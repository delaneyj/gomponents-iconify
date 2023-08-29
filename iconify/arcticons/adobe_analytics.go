package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeAnalytics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/><circle cx="24" cy="24" r="14" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.584 26.951l3.167-2.683l2.982 1.975l5.739-5.292l2.832 2.274l5.218-4.025M15.876 32.721v-4.827m4.062 7.007v-6.224M24 35.627V25.565m4.062 9.336v-9.093m4.062 6.503v-8.572"/>`),
		g.Group(children),
	)
}