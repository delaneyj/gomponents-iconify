package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1ZM12 5.828l-6 6V20h12v-8.172l-6-6Zm0 12.171a22.972 22.972 0 0 0-.692-.582l-.047-.038c-1.157-.944-2.6-2.119-2.6-3.58A1.8 1.8 0 0 1 10.5 12a2.008 2.008 0 0 1 1.5.667A2.009 2.009 0 0 1 13.5 12a1.8 1.8 0 0 1 1.835 1.8c0 1.466-1.452 2.649-2.618 3.6l-.057.047c-.237.194-.461.377-.661.554l.001-.002Z"/>`),
		g.Group(children),
	)
}