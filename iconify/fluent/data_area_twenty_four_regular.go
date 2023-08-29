package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataAreaTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3.75a.75.75 0 0 1 1.5 0v6.258l3.65-1.922a.75.75 0 0 1 .73.018l3.82 2.246l5.6-4.2a.75.75 0 0 1 1.2.6V19.5h.75a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75v-8.986a.76.76 0 0 1 0-.027V3.75Zm1.5 7.953V19.5H18V8.25l-4.8 3.6a.75.75 0 0 1-.83.046L8.48 9.608L4.5 11.703Z"/>`),
		g.Group(children),
	)
}