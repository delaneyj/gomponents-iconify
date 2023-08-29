package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowRestoreSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm5 2v4H9v10h11v-4h3V9zm2 2h7v1h-7zm0 3h7v3h-7zm-3 1h1v1.031h-1zm0 3.031h1V19h6v2h-7z"/>`),
		g.Group(children),
	)
}