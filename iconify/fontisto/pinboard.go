package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.352 14.585L8.843 19.2l.72-4.062L3.428 7.57L0 7.752L7.58-.001v2.953l7.214 6.647l4.513-1.105l-4.689 4.982L24 23.999z"/>`),
		g.Group(children),
	)
}