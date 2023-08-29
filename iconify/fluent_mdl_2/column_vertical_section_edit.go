package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnVerticalSectionEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1848 896q42 0 78 15t64 42t42 63t16 78q0 39-15 76t-43 65l-717 719l-377 94l94-377l717-718q28-28 65-42t76-15zm72 198q0-31-20-50t-52-20q-14 0-27 4t-23 15l-692 694l-34 135l135-34l692-693q21-21 21-51zm-512-838h640v556q-57-39-128-53V384h-384v532l-128 128V256zM256 384V256h256v128H256zm640-128v128H640V256h256zM256 1792v-128h256v128H256zm384 0v-128h200l-32 128H640zM0 896V640h128v256H0zm0-384V256h128v256H0zm0 1152v-256h128v256H0zm0-384v-256h128v256H0zm1152-512v256h-128V768h128zm0-384v256h-128V384h128zm0 768v128l-128 128v-256h128z"/>`),
		g.Group(children),
	)
}