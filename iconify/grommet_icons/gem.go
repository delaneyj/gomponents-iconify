package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 1h-.414l-.293.293l-3 3L3 4.586v14.828l.293.293l3 3l.293.293h10.828l.293-.293l3-3l.293-.293V4.586l-.293-.293l-3-3L17.414 1H7ZM5 6v12h1V6H5Zm3 15h8v-1H8v1Zm11-3V6h-1v12h1ZM16 3H8v1h8V3Zm0 3v12H8V6h8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}