package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Servers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 19h1v-1H7v1Zm11 0h1v-1h-1v1ZM1 23h11V1H1v22Zm11 0h11V1H12v22ZM4 5h5h-5Zm11 0h5h-5ZM4 9h5h-5Zm11 0h5h-5ZM4 13h5h-5Zm11 0h5h-5Z"/>`),
		g.Group(children),
	)
}