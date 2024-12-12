// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package json

import "fmt"

type Operation string

const (
    OperationApplied   Operation = "apply"
    OperationDestroyed Operation = "destroy"
    OperationPlanned   Operation = "plan"
)

type ChangeSummary struct {
    Add       int       `json:"add"`
    Change    int       `json:"change"`
    Import    int       `json:"import"`
    Remove    int       `json:"remove"`
    Forget    int       `json:"forget"`
    Operation Operation `json:"operation"`
}

// The summary strings for apply and plan are accidentally a public interface
// used by Terraform Cloud and Terraform Enterprise, so the exact formats of
// these strings are important.
func (cs *ChangeSummary) String() string {
    switch cs.Operation {
    case OperationApplied:
        if cs.Import > 0 {
            return fmt.Sprintf("Apply complete! Resources: %d imported, %d added, %d changed, %d destroyed, %d forgotten.", cs.Import, cs.Add, cs.Change, cs.Remove, cs.Forget)
        }
        return fmt.Sprintf("Apply complete! Resources: %d added, %d changed, %d destroyed, %d forgotten.", cs.Add, cs.Change, cs.Remove, cs.Forget)
    case OperationDestroyed:
        return fmt.Sprintf("Destroy complete! Resources: %d destroyed, %d forgotten.", cs.Remove, cs.Forget)
    case OperationPlanned:
        if cs.Import > 0 {
            return fmt.Sprintf("Plan: %d to import, %d to add, %d to change, %d to destroy, %d to forget.", cs.Import, cs.Add, cs.Change, cs.Remove, cs.Forget)
        }
        return fmt.Sprintf("Plan: %d to add, %d to change, %d to destroy, %d to forget.", cs.Add, cs.Change, cs.Remove, cs.Forget)
    default:
        return fmt.Sprintf("%s: %d add, %d change, %d destroy, %d forget", cs.Operation, cs.Add, cs.Change, cs.Remove, cs.Forget)
    }
}
