package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QqTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.791 45.5a5.839 5.839 0 0 0 0-11.678c-.047 0-.092.013-.14.015a6.842 6.842 0 0 0-13.656-.303A6.074 6.074 0 1 0 24.531 45.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.609 34.548A21.503 21.503 0 1 0 23.87 45.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.66 27.51a11.344 11.344 0 1 0-14.163 7.323"/>`),
		g.Group(children),
	)
}