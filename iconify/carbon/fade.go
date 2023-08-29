package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.24 25.14L7 26.67a13.79 13.79 0 0 0 4.18 2.44l.69-1.87a12 12 0 0 1-3.63-2.1zM4.19 18l-2 .41A14.09 14.09 0 0 0 3.86 23l1.73-1a12.44 12.44 0 0 1-1.4-4zm7.63-13.24l-.69-1.87A13.79 13.79 0 0 0 7 5.33l1.24 1.53a12 12 0 0 1 3.58-2.1zM5.59 10L3.86 9a14.37 14.37 0 0 0-1.64 4.59l2 .34A12.05 12.05 0 0 1 5.59 10zM16 2v2a12 12 0 0 1 0 24v2a14 14 0 0 0 0-28z"/>`),
		g.Group(children),
	)
}