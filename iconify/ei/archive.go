package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M42 20h-2v-5c0-.6-.4-1-1-1H11c-.6 0-1 .4-1 1v5H8v-5c0-1.7 1.3-3 3-3h28c1.7 0 3 1.3 3 3v5zm-5 20H13c-1.7 0-3-1.3-3-3V20h2v17c0 .6.4 1 1 1h24c.6 0 1-.4 1-1V20h2v17c0 1.7-1.3 3-3 3z"/><path fill="currentColor" d="M29 26h-8c-.6 0-1-.4-1-1s.4-1 1-1h8c.6 0 1 .4 1 1s-.4 1-1 1zM8 18h34v2H8z"/>`),
		g.Group(children),
	)
}