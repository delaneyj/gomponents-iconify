package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellSimpleZ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 224a8 8 0 0 1-8 8H96a8 8 0 1 1 0-16h64a8 8 0 0 1 8 8Zm-24-88h-17l23.7-35.56A8 8 0 0 0 144 88h-32a8 8 0 0 0 0 16h17.05l-23.7 35.56A8 8 0 0 0 112 152h32a8 8 0 0 0 0-16Zm77.84 56a15.8 15.8 0 0 1-13.84 8H48a16 16 0 0 1-13.8-24.06C39.75 166.38 48 139.34 48 104a80 80 0 1 1 160 0c0 35.33 8.26 62.38 13.81 71.94a15.89 15.89 0 0 1 .03 16.06ZM208 184c-7.73-13.27-16-43.95-16-80a64 64 0 1 0-128 0c0 36.06-8.28 66.74-16 80Z"/>`),
		g.Group(children),
	)
}