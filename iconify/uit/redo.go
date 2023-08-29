package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 11.5a.5.5 0 0 0-.5.5a9.02 9.02 0 1 1-1.502-5H16.5a.5.5 0 0 0 0 1h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-1 0v2.497A9.994 9.994 0 0 0 12.025 2C6.502 1.993 2.02 6.465 2.013 11.987C2.006 17.51 6.477 21.993 12 22c5.52-.006 9.994-4.48 10-10a.5.5 0 0 0-.5-.5z"/>`),
		g.Group(children),
	)
}