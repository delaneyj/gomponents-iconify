package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.787.693l5.523 5.522l-6.365 7.774l2.188 2.188l-5.659 5.659l-4.95-4.95L2.16 23.25L.746 21.836l6.364-6.364l-4.95-4.95l5.66-5.659l2.188 2.189l7.78-6.359Zm-4.313 18.314l2.83-2.83l-2.054-2.054l6.365-7.774l-2.963-2.962l-7.779 6.358L7.82 7.692l-2.83 2.83l8.485 8.485Z"/>`),
		g.Group(children),
	)
}