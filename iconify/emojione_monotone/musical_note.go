package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicalNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M25.386 2v36.721c-1.249-.406-3.727-.629-5.163-.629c-13.631 0-13.631 16.59 0 16.59c5.856 0 11.715-3.715 11.715-8.295V25.781L47.448 31v15.037c-1.249-.404-3.727-.629-5.163-.629c-13.631 0-13.631 16.592 0 16.592C48.142 62 54 58.287 54 53.705V11.442L25.386 2m22.062 21.518l-15.511-5.309v-6.291l15.511 5.367v6.233"/>`),
		g.Group(children),
	)
}