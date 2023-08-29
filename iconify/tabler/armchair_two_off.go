package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArmchairTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 10V6a3 3 0 0 1 .128-.869m2.038-2.013C7.43 3.04 7.71 3 8 3h8a3 3 0 0 1 3 3v4"/><path d="M16.124 12.145a3 3 0 1 1 3.756 3.724M19 19H5v-3a3 3 0 1 1 3-3v2m0-3h4m-5 7v2m10-2v2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}