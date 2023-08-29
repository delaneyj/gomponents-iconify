package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.21 40.09L10.28 39a2.57 2.57 0 0 1-1.78-1.79l-1.1-3.93L31.21 9.47l6.79 6.8Zm-5.05-1.76l25.46-25.46"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.219 6.816l1.464 1.464a3.78 3.78 0 0 1 0 5.346l-2.673 2.673h0l-6.81-6.81h0l2.673-2.673a3.78 3.78 0 0 1 5.346 0ZM7.4 33.28l-2.57 8.16a1 1 0 0 0 1.22 1.22l8.16-2.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.8 42.42a1.22 1.22 0 1 0-1.73-1.73"/>`),
		g.Group(children),
	)
}