package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vacation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m11 14.999l-6 1c1.63-7.514 6.364-9.993 11-10c2.997-.005 5.952 1.026 8 2c2.048-.974 5-2 8-2c4.611 0 9.37 2.486 11 10l-6-1c.559 2.1 1.788 5.792 0 9c-2.98-2.673-9.87-6.709-13-9c-3.13 2.291-10.02 6.327-13 9c-1.788-3.207-.559-6.9 0-9Z"/><path d="M24 15c-.755 3.889-1.811 13.533 0 21"/><path d="M12 42h24c-4.787-4.585-7-5.995-12-6c-5-.005-10.108 3.382-12 6Z"/></g>`),
		g.Group(children),
	)
}