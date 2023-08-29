package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VersionControlPush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 0v128H256V0h1408zm-640 603v554q83 11 153 50t122 99t80 135t29 159q0 93-35 174t-96 143t-142 96t-175 35q-93 0-174-35t-143-96t-96-142t-35-175q0-83 29-158t80-135t121-99t154-51V603L557 941l-90-90l493-493l493 493l-90 90l-339-338zm256 997q0-66-25-124t-69-101t-102-69t-124-26q-66 0-124 25t-102 69t-69 102t-25 124q0 66 25 124t68 102t102 69t125 25q66 0 124-25t101-68t69-102t26-125z"/>`),
		g.Group(children),
	)
}