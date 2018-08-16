// Copyright 2018 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package base

import (
	"fmt"

	"github.com/hexya-erp/hexya/hexya/actions"
	"github.com/hexya-erp/hexya/hexya/models"
	"github.com/hexya-erp/hexya/hexya/models/types"
	"github.com/hexya-erp/hexya/pool/h"
)

func init() {
	h.Translation().DeclareModel()
	h.Translation().Methods().TranslateFields().DeclareMethod(
		`TranslateFields opens the translation window for the given field`,
		func(rs h.TranslationSet, modelName string, id int64, fieldName models.FieldName) *actions.Action {
			fi := models.Registry.MustGet(modelName).FieldsGet(fieldName)[fieldName.String()]
			model := fmt.Sprintf("%sHexya%s", modelName, fi.Name)
			return &actions.Action{
				Name:     rs.T("Translate"),
				Type:     actions.ActionActWindow,
				Model:    model,
				ViewMode: "list",
				Domain:   fmt.Sprintf("[('record_id', '=', %d)]", id),
				Context:  types.NewContext().WithKey("default_record_id", id),
			}
		})
}
