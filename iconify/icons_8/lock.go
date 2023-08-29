package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3c-3.845 0-7 3.155-7 7v3H6v16h20V13h-3v-3c0-3.845-3.155-7-7-7zm0 2c2.755 0 5 2.245 5 5v3H11v-3c0-2.755 2.245-5 5-5zM8 15h16v12H8V15z"/>`),
		g.Group(children),
	)
}