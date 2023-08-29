package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrillDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m10 6l1.414-1.414L15 8.172V0h2v8.172l3.586-3.586L22 6l-6 6l-6-6z"/><path fill="currentColor" d="M22 16a5.981 5.981 0 0 0-1.757-4.243L16 16l-4.243-4.243A6 6 0 1 0 22 16Z"/><path fill="currentColor" d="M30 16a13.958 13.958 0 0 0-4.105-9.895l-1.414 1.414a12 12 0 1 1-16.962 0L6.105 6.105A13.997 13.997 0 1 0 30 16Z"/>`),
		g.Group(children),
	)
}