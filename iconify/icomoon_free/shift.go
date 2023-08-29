package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.5 14h-5a.5.5 0 0 1-.5-.5V8H3a.5.5 0 0 1-.354-.854l5-5a.5.5 0 0 1 .707 0l5 5a.499.499 0 0 1-.354.854h-2v5.5a.5.5 0 0 1-.5.5zM6 13h4V7.5a.5.5 0 0 1 .5-.5h1.293L8 3.207L4.207 7H5.5a.5.5 0 0 1 .5.5V13z"/>`),
		g.Group(children),
	)
}