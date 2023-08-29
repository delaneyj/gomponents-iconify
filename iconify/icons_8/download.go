package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 4v16.563L9.72 15.28l-1.44 1.44l7 7l.72.686l.72-.687l7-7l-1.44-1.44L17 20.564V4h-2zM7 27v2h18v-2H7z"/>`),
		g.Group(children),
	)
}