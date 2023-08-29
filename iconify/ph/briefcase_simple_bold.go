package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M216.008 56H180v-8a28.031 28.031 0 0 0-28-28h-48a28.031 28.031 0 0 0-28 28v8H40.008a20.022 20.022 0 0 0-20 20v128a20.022 20.022 0 0 0 20 20h176a20.022 20.022 0 0 0 20-20V76a20.022 20.022 0 0 0-20-20zM100 48a4.004 4.004 0 0 1 4-4h48a4.004 4.004 0 0 1 4 4v8h-56zm112.008 152h-168V80h168z" fill="currentColor"/>`),
		g.Group(children),
	)
}