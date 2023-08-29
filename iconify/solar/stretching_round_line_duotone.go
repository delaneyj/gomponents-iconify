package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StretchingRoundLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="14.5" cy="4.5" r="2.5"/><path stroke-linecap="round" d="m7.948 13.435l-.025-.024c-1.042-1.007-.237-2.626.67-3.268c.907-.643 4.752-1.643 4.752 3.291C13.345 18.13 9.695 22 5 22" opacity=".5"/><path stroke-linecap="round" d="M19 21.996V18.05c0-1.776-1.605-3.129-3.373-2.844"/></g>`),
		g.Group(children),
	)
}