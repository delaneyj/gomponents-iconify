package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b1cc33" d="m61.5 23.3l-8.013-8.013l-25.71 25.71l-9.26-9.26l-8.013 8.013l17.42 17.44z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M10.5 39.76L27.92 57.2L61.5 23.31l-8.013-8.013l-25.71 25.71l-9.26-9.26z"/>`),
		g.Group(children),
	)
}