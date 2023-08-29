package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreetView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4c-2.2 0-4 1.8-4 4c0 1.113.477 2.117 1.219 2.844A5.036 5.036 0 0 0 11 15v4.625l2 1V25h6v-4.375l2-1V15a5.036 5.036 0 0 0-2.219-4.156C19.523 10.117 20 9.114 20 8c0-2.2-1.8-4-4-4zm0 2c1.117 0 2 .883 2 2s-.883 2-2 2s-2-.883-2-2s.883-2 2-2zm0 6c1.668 0 3 1.332 3 3v3.375l-2 1V23h-2v-3.625l-2-1V15c0-1.668 1.332-3 3-3zm-7 6.875c-2.918.816-5 2.2-5 4.125c0 3.281 6.035 5 12 5s12-1.719 12-5c0-1.926-2.082-3.309-5-4.125v2.094c1.902.613 3 1.406 3 2.031c0 1.195-3.988 3-10 3c-6.012 0-10-1.805-10-3c0-.625 1.098-1.418 3-2.031z"/>`),
		g.Group(children),
	)
}