package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M227.45 174a4 4 0 0 1-1.44 5.48l-96 56a4 4 0 0 1-4 0l-96-56a4 4 0 0 1 4-6.92l94 54.83l94-54.83a4 4 0 0 1 5.44 1.44ZM222 124.54l-94 54.83l-94-54.83a4 4 0 0 0-4 6.92l96 56a4 4 0 0 0 4 0l96-56a4 4 0 0 0-4-6.92ZM28 80a4 4 0 0 1 2-3.46l96-56a4 4 0 0 1 4 0l96 56a4 4 0 0 1 0 6.92l-96 56a4 4 0 0 1-4 0l-96-56A4 4 0 0 1 28 80Zm11.94 0L128 131.37L216.06 80L128 28.63Z"/>`),
		g.Group(children),
	)
}