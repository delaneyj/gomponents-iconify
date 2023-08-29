package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egyptpyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 384L672 160l96-96l256 320H832zM0 640L480 0l480 640H0z"/>`),
		g.Group(children),
	)
}