package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globeone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="22.542" cy="25.543" r="7.982" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.772 33.66A19.956 19.956 0 1 1 14.108 7.457m16.231-.283a19.956 19.956 0 0 1 10.433 10.253"/><circle cx="22.344" cy="5.826" r="3.326" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="42.089" cy="25.543" r="3.326" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}