package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellZ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 144a8 8 0 0 1-8 8h-32a8 8 0 0 1-6.65-12.44l23.7-35.56H112a8 8 0 0 1 0-16h32a8 8 0 0 1 6.65 12.44L127 136h17a8 8 0 0 1 8 8Zm69.84 48a15.8 15.8 0 0 1-13.84 8h-40.81a40 40 0 0 1-78.38 0H48a16 16 0 0 1-13.8-24.06C39.75 166.38 48 139.34 48 104a80 80 0 1 1 160 0c0 35.33 8.26 62.38 13.81 71.94a15.89 15.89 0 0 1 .03 16.06Zm-71.22 8h-45.24a24 24 0 0 0 45.24 0ZM208 184c-7.73-13.27-16-43.95-16-80a64 64 0 1 0-128 0c0 36.06-8.28 66.74-16 80Z"/>`),
		g.Group(children),
	)
}