package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 19c3.038 0 5.5-2.429 5.5-6.714C16 9.429 14.167 6.333 10.5 3C6.833 6.333 5 9.429 5 12.286C5 16.57 7.462 19 10.5 19Z"/><path d="M10.5 19c1.519 0 2.75-1.214 2.75-3.357c0-1.429-.917-2.976-2.75-4.643c-1.833 1.667-2.75 3.214-2.75 4.643C7.75 17.786 8.981 19 10.5 19Z"/></g>`),
		g.Group(children),
	)
}