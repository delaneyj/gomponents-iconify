package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M38.748 15.743a.5.5 0 0 0-.45-.345l-13.146-.657l-4.685-12.298c-.146-.389-.787-.389-.934 0l-4.686 12.298l-13.145.657a.5.5 0 0 0-.288.889l10.247 8.257l-3.436 12.704c-.055.2.021.413.189.535a.504.504 0 0 0 .567.015L20 30.602l11.019 7.196a.503.503 0 0 0 .568-.015a.498.498 0 0 0 .188-.535l-3.437-12.704l10.247-8.257a.5.5 0 0 0 .163-.544z"/>`),
		g.Group(children),
	)
}