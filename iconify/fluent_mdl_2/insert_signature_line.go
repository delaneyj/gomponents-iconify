package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertSignatureLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1664 1532l128-128v644H128V0h1115l499 499q-35 11-60 23t-48 28t-42 36t-44 44l-10 10h-386V128H256v1792h1408v-388zM1280 512h293l-293-293v293zm568 128q42 0 78 15t64 42t42 63t16 78q0 39-15 76t-43 65l-717 719q-7 2-37 9t-71 18t-89 22t-86 22t-66 16t-28 7H384v-128h544l62-249l717-718q28-28 65-42t76-15zm51 249q21-21 21-51q0-31-20-50t-52-20q-14 0-27 4t-23 15l-692 694l-34 135l135-34l692-693z"/>`),
		g.Group(children),
	)
}