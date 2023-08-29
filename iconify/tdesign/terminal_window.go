package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3H1v18h22V3ZM3 10h18v9H3v-9Zm0-2V5h18v3H3Zm4 9H5v-5h2v5Z"/>`),
		g.Group(children),
	)
}