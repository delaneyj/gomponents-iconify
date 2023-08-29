package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHtmlFive0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M37.804 5H10.196a2 2 0 0 0-1.991 2.187l2.688 28.666a2 2 0 0 0 1.153 1.63l11.116 5.13a2 2 0 0 0 1.676 0l11.116-5.13a2 2 0 0 0 1.154-1.63l2.687-28.666A2 2 0 0 0 37.804 5Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M32 12H16l1 9h14l-1 11l-6 3l-6-3l-.5-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHtmlFive0)"/>`),
		g.Group(children),
	)
}