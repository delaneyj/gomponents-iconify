package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFiveSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Zm7.25-2a.5.5 0 0 0 0-1H7.001a.5.5 0 0 0-.496.438l-.25 2A.5.5 0 0 0 6.758 8h.053l.145-.002A55.282 55.282 0 0 1 7.905 8a7.236 7.236 0 0 1 .25.01l.018.002h.006a1.25 1.25 0 1 1-1.25 1.876a.5.5 0 1 0-.859.51a2.25 2.25 0 1 0 2.248-3.377c-.161-.022-.582-.025-.906-.025h-.094L7.443 6H9.25Z"/>`),
		g.Group(children),
	)
}