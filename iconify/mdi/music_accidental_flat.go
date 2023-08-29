package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicAccidentalFlat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 19c4.86-2.74 7-5.09 7-7c0-1.41-.71-3-3-3c-.7 0-1.39.28-2 .67V5h-2m2 10.38v-3.12c.62-.67 1.45-1.26 2-1.26c.59 0 1 .07 1 1c0 .15-.1 1.3-3 3.38Z"/>`),
		g.Group(children),
	)
}