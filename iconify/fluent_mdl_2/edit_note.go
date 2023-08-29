package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 1920h412l-32 128H256V128h128V0h128v128h256V0h128v128h256V0h128v128h256V0h128v128h128v648q-34 5-66 17t-62 31V256H384v1664zM1848 896q42 0 78 15t64 41t42 63t16 79q0 39-15 76t-43 65l-717 717l-377 94l94-377l717-716q29-29 65-43t76-14zm51 249q21-21 21-51q0-31-20-50t-52-20q-14 0-27 4t-23 15l-692 692l-34 135l135-34l692-691z"/>`),
		g.Group(children),
	)
}