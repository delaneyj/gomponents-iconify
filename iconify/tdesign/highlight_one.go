package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm11-5a4.998 4.998 0 0 0-4.354 2.54l-.493.871l-1.74-.985l.492-.87A6.998 6.998 0 0 1 12 5h1v2h-1Zm-4.996 4.133v2.076h-2v-2.076h2Z"/>`),
		g.Group(children),
	)
}