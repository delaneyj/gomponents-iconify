package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixbit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="33.282" height="27.12" x="7.363" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".357"/><rect width="6.067" height="6.067" x="34.527" y="37.409" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".385"/><rect width="22.137" height="6.115" x="7.355" y="37.385" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".368"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.386 19.497h9.857v3.734h-6.571v-3.734"/><circle cx="34.458" cy="16.638" r=".75" fill="currentColor"/><circle cx="13.542" cy="16.638" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}