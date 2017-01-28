// Copyright 2016 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package methods

import (
	"fmt"
	"github.com/npiganeau/yep/pool"
)

func initPartner() {
	pool.ResPartner().ExtendMethod("NameGet", "",
		func(rs pool.ResPartnerSet) string {
			res := rs.Super()
			return fmt.Sprintf("%s (%d)", res, rs.ID())
		})
}
