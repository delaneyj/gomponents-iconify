package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarRtlTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 9.5v12.25A3.25 3.25 0 0 0 6.25 25h15.5A3.25 3.25 0 0 0 25 21.75V9.5H3Zm16.259 8.001a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm-5.255 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm5.255-5a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm-5.255 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm-5.254 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5ZM6.25 3A3.25 3.25 0 0 0 3 6.25V8h22V6.25A3.25 3.25 0 0 0 21.75 3H6.25Z"/>`),
		g.Group(children),
	)
}