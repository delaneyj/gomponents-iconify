package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroitsGpsTest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="7.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="14.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="33.855" r=".75" fill="currentColor"/><circle cx="24" cy="24" r=".75" fill="currentColor"/><circle cx="16.056" cy="16.717" r=".75" fill="currentColor"/><circle cx="17.804" cy="13.822" r=".75" fill="currentColor"/><circle cx="19.065" cy="7.472" r=".75" fill="currentColor"/><circle cx="28.023" cy="7.472" r=".75" fill="currentColor"/><circle cx="38.25" cy="14.159" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}