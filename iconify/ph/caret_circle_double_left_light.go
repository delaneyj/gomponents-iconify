package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleDoubleLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200.12 55.87A102 102 0 0 0 55.87 200.12A102 102 0 1 0 200.12 55.87Zm-8.48 135.77a90 90 0 1 1 0-127.28a90.1 90.1 0 0 1 0 127.28Zm-19.4-91.4L144.49 128l27.75 27.76a6 6 0 1 1-8.48 8.48l-32-32a6 6 0 0 1 0-8.48l32-32a6 6 0 0 1 8.48 8.48Zm-56 0L88.49 128l27.75 27.76a6 6 0 1 1-8.48 8.48l-32-32a6 6 0 0 1 0-8.48l32-32a6 6 0 0 1 8.48 8.48Z"/>`),
		g.Group(children),
	)
}