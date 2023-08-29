package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Band(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0c8.837 0 16 7.163 16 16s-7.163 16-16 16S0 24.837 0 16S7.163 0 16 0zm.086 5.25L9.25 9.1v13.671l6.836 3.929l6.757-4.007v-6.757l-6.522-3.929l-2.2 1.1l6.522 3.85l.078 4.636l-4.635 2.593l-4.715-2.672V10.2l4.715-2.593l2.2 1.179v3.693l2.2 1.257V7.45l-4.4-2.2zm-1.965 8.329v5.657l4.872-2.75l-1.65-1.022l-1.493.786V14.6l-1.729-1.021z"/>`),
		g.Group(children),
	)
}