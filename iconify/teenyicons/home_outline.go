package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5.5l.325-.38a.5.5 0 0 0-.65 0L7.5.5Zm-7 6l-.325-.38L0 6.27v.23h.5Zm5 8v.5a.5.5 0 0 0 .5-.5h-.5Zm4 0H9a.5.5 0 0 0 .5.5v-.5Zm5-8h.5v-.23l-.175-.15l-.325.38ZM1.5 15h4v-1h-4v1Zm13.325-8.88l-7-6l-.65.76l7 6l.65-.76Zm-7.65-6l-7 6l.65.76l7-6l-.65-.76ZM6 14.5v-3H5v3h1Zm3-3v3h1v-3H9Zm.5 3.5h4v-1h-4v1Zm5.5-1.5v-7h-1v7h1Zm-15-7v7h1v-7H0ZM7.5 10A1.5 1.5 0 0 1 9 11.5h1A2.5 2.5 0 0 0 7.5 9v1Zm0-1A2.5 2.5 0 0 0 5 11.5h1A1.5 1.5 0 0 1 7.5 10V9Zm6 6a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1Zm-12-1a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15v-1Z"/>`),
		g.Group(children),
	)
}