package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M30.094 11.132c-.697 2.958-3.544 5.235-6.359 5.086c-2.814-.149-4.531-2.667-3.834-5.625c.697-2.958 3.544-5.235 6.359-5.086c2.815.149 4.531 2.667 3.834 5.625Z"/><path d="m31.138 26.774l8.762-5.066c5.602-2.8 1.32-11.254-4.984-8.103l-11.34 6.572l-9.518-6.354c-6.743-3.858-11.909 5.432-5.917 8.86l6.24 4.091l-1.6 11.718c3.831 4.265 11.057 5.291 16.445 2.275l1.912-13.993Z"/></g>`),
		g.Group(children),
	)
}