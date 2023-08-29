package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AspectRatioFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 4H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3zM9 7a1 1 0 0 1 .117 1.993L9 9H7v2a1 1 0 0 1-.883.993L6 12a1 1 0 0 1-.993-.883L5 11V8a1 1 0 0 1 .883-.993L6 7h3zm9 5a1 1 0 0 1 .993.883L19 13v3a1 1 0 0 1-.883.993L18 17h-3a1 1 0 0 1-.117-1.993L15 15h2v-2a1 1 0 0 1 .883-.993L18 12z"/></g>`),
		g.Group(children),
	)
}