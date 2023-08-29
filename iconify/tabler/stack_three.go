package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2L4 6l8 4l8-4l-8-4m-8 8l8 4l8-4M4 18l8 4l8-4M4 14l8 4l8-4"/>`),
		g.Group(children),
	)
}