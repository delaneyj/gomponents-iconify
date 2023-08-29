package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M17.337 4.761v30.478h-6.759V4.761h6.759m1-1H9.578v32.478h8.759V3.761zm11.085 1v30.478h-6.759V4.761h6.759m1-1h-8.759v32.478h8.759V3.761z"/>`),
		g.Group(children),
	)
}