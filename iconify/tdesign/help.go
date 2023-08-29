package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.995 21.006l-2-.006v-2.006h2.012l-.012 2.012Zm-2-3.556v-2.38c0-1.15.518-2.11 1.137-2.871c.614-.756 1.395-1.397 2.055-1.917a3.544 3.544 0 1 0-5.534-3.968l-.334.943l-1.885-.666l.333-.943A5.546 5.546 0 0 1 17.54 7.497a5.536 5.536 0 0 1-2.116 4.357c-.655.516-1.279 1.04-1.74 1.608c-.458.562-.689 1.087-.689 1.609v2.379h-2Z"/>`),
		g.Group(children),
	)
}