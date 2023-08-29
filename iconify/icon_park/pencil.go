package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M30.9995 8.99902L38.9995 16.999"/><path d="M7.99953 31.999L35.9994 4L43.9995 11.999L15.9995 39.999L5.99951 41.999L7.99953 31.999Z"/><path d="M30.9995 8.99902L38.9995 16.999"/><path d="M8.99951 31.999L15.9995 38.999"/><path d="M12.9995 34.999L34.9995 12.999"/></g>`),
		g.Group(children),
	)
}