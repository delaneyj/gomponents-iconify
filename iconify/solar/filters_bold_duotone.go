package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiltersBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M18 8A6 6 0 1 1 6 8a6 6 0 0 1 12 0Z"/><path d="M13.58 13.79a6.002 6.002 0 0 1-7.16-3.58a6 6 0 1 0 7.16 3.58Z" opacity=".7"/><path d="M13.58 13.79c.271.684.42 1.43.42 2.21a5.985 5.985 0 0 1-2 4.472a6 6 0 1 0 5.58-10.262a6.014 6.014 0 0 1-4 3.58Z" opacity=".4"/></g>`),
		g.Group(children),
	)
}