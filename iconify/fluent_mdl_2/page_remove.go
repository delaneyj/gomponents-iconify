package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1819 1728l226 227l-90 90l-227-226l-227 227l-90-91l227-227l-227-227l90-90l227 227l227-227l90 91l-226 226zM256 1920h1024v128H128V0h1115l549 549v731h-128V640h-512V128H256v1792zM1280 512h293l-293-293v293z"/>`),
		g.Group(children),
	)
}