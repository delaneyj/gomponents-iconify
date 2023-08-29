package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Therapeer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="26.419" height="39" x="10.79" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.553"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.8 41.12l14.06-4.178a2.516 2.516 0 0 0 1.8-2.412V13.47a2.516 2.516 0 0 0-1.8-2.412L10.8 6.88m12.579 19.074v-3.908"/>`),
		g.Group(children),
	)
}