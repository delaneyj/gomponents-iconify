package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextboxAlignMiddleRotateNinetyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 5.75A2.75 2.75 0 0 0 18.25 3H5.75A2.75 2.75 0 0 0 3 5.75v12.5A2.75 2.75 0 0 0 5.75 21h12.5A2.75 2.75 0 0 0 21 18.25V5.75Zm-6 1.5v9.5a.75.75 0 0 1-1.5 0v-9.5a.75.75 0 0 1 1.5 0Zm-4 0v9.5a.75.75 0 0 1-1.5 0v-9.5a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}