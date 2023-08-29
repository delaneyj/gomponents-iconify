package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feStepBackward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStepBackward1" fill="currentColor" fill-rule="nonzero"><path id="feStepBackward2" d="M11 15V5H4a1 1 0 1 1 0-2h16a1 1 0 0 1 0 2h-7v10h2.24c.15 0 .297.042.421.12c.35.219.444.663.211.991l-3.24 4.57a.74.74 0 0 1-.21.199a.79.79 0 0 1-1.054-.198l-3.24-4.57A.685.685 0 0 1 8 15.714c0-.395.34-.715.76-.715H11Zm9-6a1 1 0 0 1 0 2h-5V9h5ZM8 9h1v2H4a1 1 0 0 1 0-2h4Z"/></g></g>`),
		g.Group(children),
	)
}