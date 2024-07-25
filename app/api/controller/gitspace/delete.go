// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gitspace

import (
	"context"
	"fmt"

	apiauth "github.com/harness/gitness/app/api/auth"
	"github.com/harness/gitness/app/auth"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"

	"github.com/rs/zerolog/log"
)

const gitspaceConfigNotFound = "Failed to find gitspace config with identifier "

func (c *Controller) Delete(
	ctx context.Context,
	session *auth.Session,
	spaceRef string,
	identifier string,
) error {
	space, err := c.spaceStore.FindByRef(ctx, spaceRef)
	if err != nil {
		return fmt.Errorf("failed to find space: %w", err)
	}
	err = apiauth.CheckGitspace(ctx, c.authorizer, session, space.Path, identifier, enum.PermissionGitspaceDelete)
	if err != nil {
		return fmt.Errorf("failed to authorize: %w", err)
	}

	gitspaceConfig, err := c.gitspaceConfigStore.FindByIdentifier(ctx, space.ID, identifier)
	if err != nil || gitspaceConfig == nil {
		log.Err(err).Msg(gitspaceConfigNotFound + identifier)
		return err
	}
	instance, _ := c.gitspaceInstanceStore.FindLatestByGitspaceConfigID(ctx, gitspaceConfig.ID, gitspaceConfig.SpaceID)
	gitspaceConfig.GitspaceInstance = instance
	gitspaceConfig.SpacePath = space.Path
	if instance != nil {
		if stopErr := c.stopRunningGitspace(ctx, gitspaceConfig); stopErr != nil {
			return stopErr
		}
	}
	gitspaceConfig.IsDeleted = true
	if err = c.gitspaceConfigStore.Update(ctx, gitspaceConfig); err != nil {
		return fmt.Errorf("failed to delete gitspace config with ID: %s %w", gitspaceConfig.Identifier, err)
	}
	return nil
}

func (c *Controller) stopRunningGitspace(
	ctx context.Context,
	config *types.GitspaceConfig,
) error {
	if instanceUpdated, err := c.orchestrator.DeleteGitspace(ctx, config); err != nil {
		return err
	} else if err = c.gitspaceInstanceStore.Update(ctx, instanceUpdated); err != nil {
		return err
	}
	return nil
}
