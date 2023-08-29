package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Woman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M32.485 24.485A11.962 11.962 0 0 0 36 16c0-6.627-5.373-12-12-12S12 9.373 12 16c0 3.314 1.343 6.314 3.515 8.485"/><path stroke-linecap="round" stroke-linejoin="round" d="m6 44l1-5l11-8l6 6l6-6l11 8l1 5"/><path d="M12.993 21.007c.013-4.11.682-7.113 2.007-9.007c1.988-2.841 3.387-2.632 4.405-2.19c1.019.443 1.618 3.334 3.319 4.168c1.7.833 6.054.936 7.545 1.939c1.49 1.002 4.9 2.867 4.05 6.051"/></g>`),
		g.Group(children),
	)
}