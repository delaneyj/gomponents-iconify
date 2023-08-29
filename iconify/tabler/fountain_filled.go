package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FountainFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M15 2a4 4 0 0 1 4 4a1 1 0 0 1-1.993.117L17 6a2 2 0 0 0-3.995-.15L13 6v9h1v-4a3 3 0 0 1 6 0a1 1 0 0 1-1.993.117L18 11a1 1 0 0 0-1.993-.117L16 11v4h5a1 1 0 0 1 .993.883L22 16v2a4 4 0 0 1-3.8 3.995L18 22H6a4 4 0 0 1-3.995-3.8L2 18v-2a1 1 0 0 1 .883-.993L3 15h5v-4a1 1 0 0 0-1.993-.117L6 11a1 1 0 0 1-2 0a3 3 0 0 1 5.995-.176L10 11v4h1V6a2 2 0 1 0-4 0a1 1 0 1 1-2 0a4 4 0 0 1 7.001-2.645A3.983 3.983 0 0 1 15 2z"/></g>`),
		g.Group(children),
	)
}