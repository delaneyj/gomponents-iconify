package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookDownloader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="38.5" cy="38.499" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.432 43.326A21.413 21.413 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24c0 3.383-.781 6.583-2.173 9.43M38.5 42.499v-8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.203 39.203l3.297 3.296l3.297-3.296M22.965 33.5V17.997a3.497 3.497 0 0 1 3.497-3.497h.007c1.717 0 2.793.51 3.53 1.462M20 21.252h7"/>`),
		g.Group(children),
	)
}