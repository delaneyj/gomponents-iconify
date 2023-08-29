package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subtitles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7h-8v10h8v-2h-6V9h6V7zM3 15V7h8v2H5v6h6v2H3v-2z"/>`),
		g.Group(children),
	)
}