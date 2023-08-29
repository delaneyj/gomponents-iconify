package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suntimesalarms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="10" height="5" x="2.7" y="5.27" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1" transform="rotate(-44.72 7.703 7.773)"/><rect width="5" height="10" x="37.58" y="2.77" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1" transform="rotate(-45.28 40.074 7.778)"/><circle cx="24" cy="24" r="18.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24h37M24 5.5V24"/>`),
		g.Group(children),
	)
}