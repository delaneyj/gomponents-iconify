package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineDashesThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M28.634 3.366a1.25 1.25 0 0 1 0 1.768l-.98.979a1.25 1.25 0 1 1-1.767-1.768l.98-.979a1.25 1.25 0 0 1 1.767 0Zm-4.896 4.896a1.25 1.25 0 0 1 0 1.768l-1.958 1.958a1.25 1.25 0 1 1-1.768-1.768l1.958-1.958a1.25 1.25 0 0 1 1.768 0Zm-5.875 5.875a1.25 1.25 0 0 1 0 1.768l-1.958 1.958a1.25 1.25 0 1 1-1.768-1.768l1.958-1.958a1.25 1.25 0 0 1 1.768 0Zm-5.875 5.875a1.25 1.25 0 0 1 0 1.768l-1.958 1.958a1.25 1.25 0 0 1-1.768-1.768l1.958-1.958a1.25 1.25 0 0 1 1.768 0Zm-5.875 5.875a1.25 1.25 0 0 1 0 1.768l-.98.979a1.25 1.25 0 0 1-1.767-1.768l.98-.98a1.25 1.25 0 0 1 1.767 0Z"/>`),
		g.Group(children),
	)
}