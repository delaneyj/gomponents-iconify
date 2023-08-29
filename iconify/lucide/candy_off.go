package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CandyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8.5 8.5l-1 1a4.95 4.95 0 0 0 7 7l1-1m-3.657-9.313A4.947 4.947 0 0 1 16.5 7.5a4.947 4.947 0 0 1 1.313 4.657M14 16.5V14m0-7.5v1.843M10 10v7.5"/><path d="m16 7l1-5l1.367.683A3 3 0 0 0 19.708 3H21v1.292a3 3 0 0 0 .317 1.341L22 7l-5 1m-9 9l-1 5l-1.367-.683A3 3 0 0 0 4.292 21H3v-1.292a3 3 0 0 0-.317-1.341L2 17l5-1M2 2l20 20"/></g>`),
		g.Group(children),
	)
}