package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0c8.837 0 16 7.163 16 16s-7.163 16-16 16S0 24.837 0 16S7.163 0 16 0zm-2.255 23.42l-.048.03l-1.915 1.108L16 27l1.963-1.138l-4.145-2.399l-.073-.042zm6.477-11.08l-5.277 3.057v6.111l5.27 3.052l.007.004l5.278-3.056v-6.11l-5.278-3.058zM9.536 16.221L6.5 17.98v3.514l3.036 1.758l3.035-1.758V17.98l-3.035-1.76zM16 5l-9.5 5.5v4.87l1.91-1.105l1.125-.652l1.126.652l2.031 1.176v-1.348l1.126-.652l5.278-3.055l1.126-.652l1.126.652l4.152 2.403v-2.29L16 5z"/>`),
		g.Group(children),
	)
}