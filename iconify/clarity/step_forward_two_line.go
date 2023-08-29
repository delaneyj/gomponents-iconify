package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepForwardTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M7.08 6.52a1.68 1.68 0 0 0 0 2.4L16.51 18l-9.39 9.08a1.7 1.7 0 0 0 2.36 2.44L21.4 18L9.48 6.47a1.69 1.69 0 0 0-2.4.05Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M26.49 5a1.7 1.7 0 0 0-1.7 1.7v22.6a1.7 1.7 0 0 0 3.4 0V6.7a1.7 1.7 0 0 0-1.7-1.7Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}