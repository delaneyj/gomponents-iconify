package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrillBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m6 22l1.414-1.414L3.828 17H12v-2H3.828l3.586-3.586L6 10l-6 6l6 6z"/><path fill="currentColor" d="M16 10a5.981 5.981 0 0 0-4.243 1.757L16 16l-4.243 4.243A6 6 0 1 0 16 10Z"/><path fill="currentColor" d="M16 2a13.958 13.958 0 0 0-9.895 4.105l1.414 1.414a12 12 0 1 1 0 16.962l-1.414 1.414A13.997 13.997 0 1 0 16 2Z"/>`),
		g.Group(children),
	)
}