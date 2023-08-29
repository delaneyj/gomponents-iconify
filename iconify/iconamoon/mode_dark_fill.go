package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeDarkFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.535 3.518a1 1 0 0 0-1.061-1.402C5.675 2.852 2 6.996 2 12c0 5.523 4.477 10 10 10s10-4.477 10-10a10.4 10.4 0 0 0-.004-.28a1 1 0 0 0-1.571-.793a6 6 0 0 1-8.89-7.409Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}