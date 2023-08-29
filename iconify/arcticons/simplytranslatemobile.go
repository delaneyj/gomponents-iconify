package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplytranslatemobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.116 19.217l3.123-9.25m0 0l3.006 9.25m-1.04-3.122h-4.048m-9.951 22.242a11.027 11.027 0 0 0 9.25 4.163h5.55a9.25 9.25 0 0 0 0-18.5h-6.012a9.25 9.25 0 0 1 0-18.5h5.55c4.162 0 6.937.925 9.25 4.162"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.794 7.35v2.313H33.48M14.519 38.337h-2.313v2.313m7.524-10.18h8.54M24 29.047v1.424m-2.135 2.134a13.983 13.983 0 0 0 4.981 5.693"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.585 38.297c3.602-1.142 5.416-4.438 6.261-7.826"/>`),
		g.Group(children),
	)
}