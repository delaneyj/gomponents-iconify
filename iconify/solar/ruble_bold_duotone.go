package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RubleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z" opacity=".5"/><path d="M9 13.25a.75.75 0 1 0 0 1.5h.25V17a.75.75 0 0 0 1.5 0v-2.25H12a.75.75 0 0 0 0-1.5h-1.25v-.5h2.75a3.25 3.25 0 0 0 0-6.5h-2.338c-.146 0-.297 0-.436.022a1.75 1.75 0 0 0-1.454 1.454c-.022.139-.022.29-.022.436v3.088H9a.75.75 0 1 0 0 1.5h.25v.5H9Zm4.5-2h-2.75V8.2l.001-.198l.002-.043a.25.25 0 0 1 .206-.206a8.208 8.208 0 0 1 .24-.003H13.5a1.75 1.75 0 1 1 0 3.5Z"/></g>`),
		g.Group(children),
	)
}