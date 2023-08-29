package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StretchingRoundBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="14.5" cy="4.5" r="2.5"/><path stroke-linecap="round" d="M19 21.996v-3.947c0-1.776-1.605-3.13-3.373-2.844m-7.679-1.77l-.025-.024c-1.042-1.007-.237-2.626.67-3.268c.907-.643 4.752-1.643 4.752 3.291a8.746 8.746 0 0 1-1.302 4.615M5 22c1.46 0 2.82-.374 4-1.032"/></g>`),
		g.Group(children),
	)
}