package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.853 7.147a.5.5 0 0 0-.707.707L11.302 12l-4.156 4.146a.5.5 0 1 0 .707.708l4.51-4.5a.5.5 0 0 0 0-.707l-4.51-4.5zM16.5 7a.5.5 0 0 0-.5.5v9a.5.5 0 1 0 1 0v-9a.5.5 0 0 0-.5-.5z"/>`),
		g.Group(children),
	)
}