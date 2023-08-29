package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m.5 5.5l-.29-.407l-.21.15V5.5h.5Zm7-5l.29-.407a.5.5 0 0 0-.58 0L7.5.5Zm7 5h.5v-.257l-.21-.15l-.29.407Zm-12 3V8H2v.5h.5Zm4 0H7V8h-.5v.5ZM1 15V5.5H0V15h1ZM.79 5.907l7-5l-.58-.814l-7 5l.58.814Zm6.42-5l7 5l.58-.814l-7-5l-.58.814ZM14 5.5V15h1V5.5h-1ZM3 15V8.5H2V15h1Zm-.5-6h4V8h-4v1ZM6 8.5V15h1V8.5H6ZM5 12h1.5v-1H5v1Zm6-4v4h1V8h-1Zm2-6v2.5h1V2h-1Z"/>`),
		g.Group(children),
	)
}