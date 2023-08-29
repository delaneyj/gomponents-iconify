package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13v9H2v-9h20Zm-2 2H4v5h16v-5Zm2-13v9H2V2h20Zm-2 2H4v5h16V4ZM7.504 16.5v2.004H5.5V16.5h2.004Zm0-11v2.004H5.5V5.5h2.004Z"/>`),
		g.Group(children),
	)
}