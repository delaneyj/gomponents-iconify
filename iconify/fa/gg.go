package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m672 576l384 384l-384 384L0 672L672 0l168 168l-96 96l-72-72l-480 480l480 480l193-193l-289-287zM1248 0l672 672l-672 672l-168-168l96-96l72 72l480-480l-480-480l-193 193l289 287l-96 96l-384-384z"/>`),
		g.Group(children),
	)
}