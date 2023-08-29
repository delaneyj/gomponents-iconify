package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Npm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12.534V9.871h-1.334v2.666zM24 7.2H0v8h6.666v1.334H12v-1.333h12zM6.666 8.534v5.337H5.333v-4H3.999v4H1.333V8.537zm6.667 0v5.337h-2.666v1.334H8.001V8.539zm9.333 0v5.337h-1.333v-4h-1.334v4h-1.334v-4h-1.333v4h-2.667V8.537z"/>`),
		g.Group(children),
	)
}