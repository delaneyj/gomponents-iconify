package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2h-5v20h5c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1zM2 21c0 .6.4 1 1 1h11v-9H2v8zM2 3v8h12V2H3c-.6 0-1 .4-1 1z"/>`),
		g.Group(children),
	)
}