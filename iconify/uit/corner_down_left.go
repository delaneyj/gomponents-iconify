package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 3a.5.5 0 0 0-.5.5v10a2.502 2.502 0 0 1-2.5 2.5H6.707l3.646-3.646a.5.5 0 0 0-.707-.707l-4.5 4.5a.499.499 0 0 0-.145.35L5 16.5a.5.5 0 0 0 .146.354l4.5 4.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708L6.708 17H15.5a3.504 3.504 0 0 0 3.5-3.5v-10a.5.5 0 0 0-.5-.5z"/>`),
		g.Group(children),
	)
}