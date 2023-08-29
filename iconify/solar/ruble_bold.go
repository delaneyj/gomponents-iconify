package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RubleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13.5 11.25h-2.75V8.2l.001-.198l.002-.043a.25.25 0 0 1 .206-.206a8.208 8.208 0 0 1 .24-.003H13.5a1.75 1.75 0 1 1 0 3.5Z"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10ZM8.25 14a.75.75 0 0 1 .75-.75h.25v-.5H9a.75.75 0 0 1 0-1.5h.25V8.162c0-.146 0-.297.022-.436a1.75 1.75 0 0 1 1.454-1.454c.139-.022.29-.022.435-.022H13.5a3.25 3.25 0 0 1 0 6.5h-2.75v.5H12a.75.75 0 0 1 0 1.5h-1.25V17a.75.75 0 0 1-1.5 0v-2.25H9a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}