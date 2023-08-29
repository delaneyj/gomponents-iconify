package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 4a4 4 0 0 1 0 8H8a4.992 4.992 0 0 0 2-4a4.992 4.992 0 0 0-2-4h3zm-6 8a4 4 0 1 1 0-8a4 4 0 0 1 0 8zM0 8a5 5 0 0 0 5 5h6a5 5 0 0 0 0-10H5a5 5 0 0 0-5 5z"/>`),
		g.Group(children),
	)
}