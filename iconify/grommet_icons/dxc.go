package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dxc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12 14l4 7H8l4-7Zm0-4L8 3h8l-4 7ZM2 18H0V6h2a6 6 0 1 1 0 12Zm20 0a6 6 0 1 1 0-12h2v12h-2Z"/>`),
		g.Group(children),
	)
}