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

package keywordsearch

import (
	"context"
	"fmt"

	"github.com/harness/gitness/types"
)

type LocalIndexSearcher struct {
}

func NewLocalIndexSearcher() *LocalIndexSearcher {
	return &LocalIndexSearcher{}
}

func (s *LocalIndexSearcher) Search(
	_ context.Context,
	_ []int64,
	_ string,
	_ int,
) (types.SearchResult, error) {
	return types.SearchResult{}, fmt.Errorf("not implemented")
}

func (s *LocalIndexSearcher) Index(_ context.Context, _ *types.Repository) error {
	return fmt.Errorf("not implemented")
}