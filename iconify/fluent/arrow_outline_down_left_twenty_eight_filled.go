package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOutlineDownLeftTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.544 2.659a2.25 2.25 0 0 0-3.182 0L8.76 9.262l-1.6-1.6C5.81 6.313 3.5 7.149 3.328 9.049L2.009 23.546a2.25 2.25 0 0 0 2.445 2.445l14.5-1.318c1.9-.173 2.736-2.483 1.387-3.832l-1.602-1.601l6.602-6.602a2.25 2.25 0 0 0 0-3.182L18.544 2.66Z"/>`),
		g.Group(children),
	)
}