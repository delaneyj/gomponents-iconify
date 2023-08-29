package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plugins(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 8V3a1 1 0 0 0-2 0v5h-4V3a1 1 0 0 0-2 0v5H5v4a6.992 6.992 0 0 0 4 6.317V22h6v-3.683A6.994 6.994 0 0 0 19 12V8h-3z"/>`),
		g.Group(children),
	)
}