package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AxisChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 5H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2ZM4 29V7h28v22Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="m6.007 26.731l21.723-.001c.76-.06 1.18-.93.74-1.56a.974.974 0 0 0-.74-.43l-19.729-.004l.009-13.726c0-.78-.84-1.26-1.51-.87c-.31.17-.5.51-.5.87l.007 15.721Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}