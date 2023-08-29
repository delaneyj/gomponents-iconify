package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m6 10l2-1l7-7l-1-1l-7 7l-1 2zm-1.48 3.548c-.494-1.043-1.026-1.574-2.069-2.069l1.548-4.262l2-1.217l6-6h-3l-6 6l-3 10l10-3l6-6V4l-6 6l-1.217 2z"/>`),
		g.Group(children),
	)
}