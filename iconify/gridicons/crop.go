package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 16h-4V8a2 2 0 0 0-2-2H8V2H6v4H2v2h4v8a2 2 0 0 0 2 2h8v4h2v-4h4v-2zM8 16V8h8v8H8z"/>`),
		g.Group(children),
	)
}