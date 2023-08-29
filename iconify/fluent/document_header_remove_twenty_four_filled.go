package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentHeaderRemoveTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M17.5 12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zm-2.476 3.024a.5.5 0 0 0 0 .708l1.77 1.77l-1.767 1.766a.5.5 0 0 0 .707.707l1.766-1.767l1.77 1.77a.5.5 0 0 0 .707-.708l-1.769-1.769l1.772-1.77a.5.5 0 1 0-.708-.707l-1.77 1.77l-1.77-1.77a.5.5 0 0 0-.708 0z" fill="currentColor"/><path d="M8.75 6.5a.25.25 0 0 0-.25.25v1.5c0 .138.112.25.25.25h6.5a.25.25 0 0 0 .25-.25v-1.5a.25.25 0 0 0-.25-.25h-6.5z" fill="currentColor"/><path d="M17.745 1.996a2.25 2.25 0 0 1 2.245 2.096l.005.154v7.25a6.5 6.5 0 0 0-7.191 10.498H6.245A2.25 2.25 0 0 1 4 19.898l-.005-.154V4.246a2.25 2.25 0 0 1 2.096-2.245l.154-.005h11.5zM7 6.75v1.5C7 9.217 7.784 10 8.75 10h6.5A1.75 1.75 0 0 0 17 8.25v-1.5A1.75 1.75 0 0 0 15.25 5h-6.5A1.75 1.75 0 0 0 7 6.75z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}