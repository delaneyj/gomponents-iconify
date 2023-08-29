package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minimize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.707 3.707a1 1 0 0 0-1.414-1.414L16 6.586V5a1 1 0 1 0-2 0v4a1 1 0 0 0 1 1h4a1 1 0 1 0 0-2h-1.586l4.293-4.293zm-19.414 18a1 1 0 0 1 0-1.414L6.586 16H5a1 1 0 1 1 0-2h4a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-1.586l-4.293 4.293a1 1 0 0 1-1.414 0z"/>`),
		g.Group(children),
	)
}