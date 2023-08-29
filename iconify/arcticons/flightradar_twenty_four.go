package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightradarTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="14" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="6.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r=".75" fill="currentColor"/><circle cx="14.601" cy="18.911" r=".75" fill="currentColor"/><circle cx="13.851" cy="38.75" r=".75" fill="currentColor"/><circle cx="28.486" cy="35.47" r=".75" fill="currentColor"/><circle cx="33.449" cy="8.845" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 24L9.102 39.5"/>`),
		g.Group(children),
	)
}