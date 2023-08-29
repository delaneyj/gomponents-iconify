package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RabbitZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 44V24C6.16667 22.3334 7.2 17.1998 12 13.9998L8 9.99974C6 7.54786 7.5 2.00024 13 3.99998C15 5.00024 16 6.50024 17.5 8.00024L24 15.0002C24.6667 16.0002 26.5 17.809 26 22.4522"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 25.0005C21 25.0005 31 28.0005 34 36.0005C34.6667 38.2624 35.5 41.9995 36 43.9995"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 43.9996C18 40.1092 20.8 31.9578 32 33.4517"/><circle cx="40.5" cy="41.452" r="2" fill="#000"/></g>`),
		g.Group(children),
	)
}