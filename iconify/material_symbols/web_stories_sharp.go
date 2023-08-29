package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebStoriesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20V4h2v16h-2ZM2 22V2h13v20H2Zm19-4V6h1.5v12H21Z"/>`),
		g.Group(children),
	)
}