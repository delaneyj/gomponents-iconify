package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvenMittsMinimalisticBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.588 20.026l-3.57-3.49C2.674 15.222 2 14.564 2 13.747c0-.537.29-1.005.87-1.635c.59-.643.886-.964 1.02-1.3c.133-.336.137-.714.144-1.47L4.066 6c-.034-2.183 1.375-3.973 3.147-4c1.455-.022 2.702 1.152 3.121 2.78l.466-.456c2.562-2.505 6.716-2.505 9.278 0a6.314 6.314 0 0 1 0 9.072l-6.78 6.63C11.952 21.342 11.278 22 10.443 22c-.837 0-1.51-.658-2.855-1.974ZM5.98 12.872a.75.75 0 0 1 1.06-.012l4.283 4.187a.75.75 0 0 1-1.048 1.073l-4.283-4.187a.75.75 0 0 1-.012-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}