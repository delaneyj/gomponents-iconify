package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enquire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-miterlimit="2" d="M37 16C34.2386 16 32 13.7614 32 11C32 8.23858 34.2386 6 37 6C39.7614 6 42 8.23858 42 11C42 13.7614 39.7614 16 37 16Z"/><path fill="#2F88FF" stroke-miterlimit="2" d="M12 12C9.79086 12 8 10.2091 8 8C8 5.79086 9.79086 4 12 4C14.2091 4 16 5.79086 16 8C16 10.2091 14.2091 12 12 12Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M26 39L32 34V28C32 24.5339 34 22 37 22C40 22 42 24.5339 42 28V32.8372V42"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 33L18 28V24C18 20.5339 16 18 13 18C10 18 8 20.5339 8 24V26.8372V42"/></g>`),
		g.Group(children),
	)
}