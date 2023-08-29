package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Turingmachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 33.75l9.71 7.93V25.83L4.5 33.75zm9.71 0h19.35a9.16 9.16 0 0 0 9.16-9.16h0a9.29 9.29 0 0 0-.23-2.06m1.01-8.28l-9.71-7.93v15.85l9.71-7.92z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.79 14.25H14.44a9.16 9.16 0 0 0-9.16 9.16h0a9.29 9.29 0 0 0 .23 2.06"/>`),
		g.Group(children),
	)
}