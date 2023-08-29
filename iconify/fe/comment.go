package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feComment0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feComment1" fill="currentColor"><path id="feComment2" d="M7.268 18.732L5 21v-4.157c-1.25-1.46-2-3.319-2-5.343C3 6.806 7.03 3 12 3s9 3.806 9 8.5s-4.03 8.5-9 8.5a9.352 9.352 0 0 1-4.732-1.268Z"/></g></g>`),
		g.Group(children),
	)
}