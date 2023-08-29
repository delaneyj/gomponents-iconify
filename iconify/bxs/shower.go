package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18.33A6.78 6.78 0 0 0 19.5 15a6.73 6.73 0 0 0-1.5 3.33a1.51 1.51 0 1 0 3 0zm-10 2A6.78 6.78 0 0 0 9.5 17A6.73 6.73 0 0 0 8 20.33A1.59 1.59 0 0 0 9.5 22a1.59 1.59 0 0 0 1.5-1.67zm5 0A6.78 6.78 0 0 0 14.5 17a6.73 6.73 0 0 0-1.5 3.33A1.59 1.59 0 0 0 14.5 22a1.59 1.59 0 0 0 1.5-1.67zm-10-2A6.78 6.78 0 0 0 4.5 15A6.73 6.73 0 0 0 3 18.33A1.59 1.59 0 0 0 4.5 20A1.59 1.59 0 0 0 6 18.33zM2 12h20v2H2zm11-7.93V2h-2v2.07A8 8 0 0 0 4.07 11h15.86A8 8 0 0 0 13 4.07z"/>`),
		g.Group(children),
	)
}