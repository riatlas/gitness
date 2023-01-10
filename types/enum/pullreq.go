// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package enum

import (
	"strings"
)

// PullReqState defines pull request state.
type PullReqState string

// PullReqState enumeration.
const (
	PullReqStateOpen     PullReqState = "open"
	PullReqStateMerged   PullReqState = "merged"
	PullReqStateClosed   PullReqState = "closed"
	PullReqStateRejected PullReqState = "rejected"
)

// PullReqSort defines pull request attribute that can be used for sorting.
type PullReqSort int

// PullReqAttr enumeration.
const (
	PullReqSortNone PullReqSort = iota
	PullReqSortNumber
	PullReqSortCreated
	PullReqSortUpdated
)

// ParsePullReqSort parses the pull request attribute string
// and returns the equivalent enumeration.
func ParsePullReqSort(s string) PullReqSort {
	switch strings.ToLower(s) {
	case number:
		return PullReqSortNumber
	case created, createdAt:
		return PullReqSortCreated
	case updated, updatedAt:
		return PullReqSortUpdated
	default:
		return PullReqSortNone
	}
}

// String returns the string representation of the attribute.
func (a PullReqSort) String() string {
	switch a {
	case PullReqSortNumber:
		return number
	case PullReqSortCreated:
		return created
	case PullReqSortUpdated:
		return updated
	case PullReqSortNone:
		return ""
	default:
		return undefined
	}
}

// PullReqActivityType defines pull request activity message type.
// Essentially, the Type determines the structure of the pull request activity's Payload structure.
type PullReqActivityType string

// PullReqActivityType enumeration.
const (
	PullReqActivityTypeComment      PullReqActivityType = "comment"
	PullReqActivityTypeCodeComment  PullReqActivityType = "code-comment"
	PullReqActivityTypeTitleChange  PullReqActivityType = "title-change"
	PullReqActivityTypeReviewSubmit PullReqActivityType = "review-submit"
	PullReqActivityTypeMerge        PullReqActivityType = "merge"
)

func GetAllPullReqActivityTypes() []PullReqActivityType {
	return []PullReqActivityType{
		PullReqActivityTypeComment,
		PullReqActivityTypeCodeComment,
		PullReqActivityTypeTitleChange,
		PullReqActivityTypeReviewSubmit,
	}
}

var rawPullReqActivityTypes = toSortedStrings(GetAllPullReqActivityTypes())

// ParsePullReqActivityType parses the pull request activity type.
func ParsePullReqActivityType(s string) (PullReqActivityType, bool) {
	if existsInSortedSlice(rawPullReqActivityTypes, s) {
		return PullReqActivityType(s), true
	}
	return "", false
}

// PullReqActivityKind defines kind of pull request activity system message.
// Kind defines the source of the pull request activity entry:
// Whether it's generated by the system, it's a user comment or a part of code review.
type PullReqActivityKind string

// PullReqActivityKind enumeration.
const (
	PullReqActivityKindSystem      PullReqActivityKind = "system"
	PullReqActivityKindComment     PullReqActivityKind = "comment"
	PullReqActivityKindCodeComment PullReqActivityKind = "code"
)

func GetAllPullReqActivityKinds() []PullReqActivityKind {
	return []PullReqActivityKind{
		PullReqActivityKindSystem,
		PullReqActivityKindComment,
		PullReqActivityKindCodeComment,
	}
}

var rawPullReqActivityKinds = toSortedStrings(GetAllPullReqActivityKinds())

// ParsePullReqActivityKind parses the pull request activity type.
func ParsePullReqActivityKind(s string) (PullReqActivityKind, bool) {
	if existsInSortedSlice(rawPullReqActivityKinds, s) {
		return PullReqActivityKind(s), true
	}
	return "", false
}

// PullReqReviewDecision defines state of a pull request review.
type PullReqReviewDecision string

// PullReqReviewDecision enumeration.
const (
	PullReqReviewDecisionPending   PullReqReviewDecision = "pending"
	PullReqReviewDecisionReviewed  PullReqReviewDecision = "reviewed"
	PullReqReviewDecisionApproved  PullReqReviewDecision = "approved"
	PullReqReviewDecisionChangeReq PullReqReviewDecision = "changereq"
)

func GetAllPullReqReviewDecisions() []PullReqReviewDecision {
	return []PullReqReviewDecision{
		PullReqReviewDecisionPending,
		PullReqReviewDecisionReviewed,
		PullReqReviewDecisionApproved,
		PullReqReviewDecisionChangeReq,
	}
}

var rawPullReqReviewDecisions = toSortedStrings(GetAllPullReqReviewDecisions())

// ParsePullReqReviewDecision parses the pull request review state type.
func ParsePullReqReviewDecision(s string) (PullReqReviewDecision, bool) {
	if existsInSortedSlice(rawPullReqReviewDecisions, s) {
		return PullReqReviewDecision(s), true
	}
	return "", false
}

// PullReqReviewerType defines type of a pull request reviewer.
type PullReqReviewerType string

// PullReqReviewerType enumeration.
const (
	PullReqReviewerTypeRequested    PullReqReviewerType = "requested"
	PullReqReviewerTypeAssigned     PullReqReviewerType = "assigned"
	PullReqReviewerTypeSelfAssigned PullReqReviewerType = "self_assigned"
)

func GetAllPullReqReviewerTypes() []PullReqReviewerType {
	return []PullReqReviewerType{
		PullReqReviewerTypeRequested,
		PullReqReviewerTypeAssigned,
		PullReqReviewerTypeSelfAssigned,
	}
}

var rawPullReqReviewerTypes = toSortedStrings(GetAllPullReqReviewerTypes())

// ParsePullReqReviewerType parses the pull request reviewer type.
func ParsePullReqReviewerType(s string) (PullReqReviewerType, bool) {
	if existsInSortedSlice(rawPullReqReviewerTypes, s) {
		return PullReqReviewerType(s), true
	}
	return "", false
}

// MergeMethod represents the approach to merge commits into base branch.
type MergeMethod string

const (
	// MergeMethodMerge create merge commit.
	MergeMethodMerge MergeMethod = "merge"
	// MergeMethodSquash squash commits into single commit before merging.
	MergeMethodSquash MergeMethod = "squash"
	// MergeMethodRebase rebase before merging.
	MergeMethodRebase MergeMethod = "rebase"
)
