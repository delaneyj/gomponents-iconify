package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.947.165a1 1 0 0 0-.894 0l-6.5 3.25A1 1 0 0 0 0 4.309V12a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V4.309a1 1 0 0 0-.553-.894l-6.5-3.25Zm5.622 3.928L7.5 1.06L1.431 4.093L7.5 7.291l6.069-3.198ZM1 4.883V12h13V4.884L7.71 8.198a.45.45 0 0 1-.42 0L1 4.884Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}