package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroidifyAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="16.427" cy="12.512" r="3.776" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.551" cy="12.512" r="3.776" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="33.785" height="15.085" x="7.107" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.135"/><rect width="33.785" height="18.673" x="7.107" y="23.827" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.638"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.709 16.92a2.996 2.996 0 0 0 4.726 0"/>`),
		g.Group(children),
	)
}