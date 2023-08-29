package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopSelected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1016 1592l72 72H128q-25 0-48-10t-41-29t-28-41t-11-48q0-16 3-36t10-39t16-36t22-30l205-206V384h1408v805l-91 91H347l-203 202q-3 3-6 10t-5 16t-4 16t-1 12h945l-57 56zm-632-440h1152V512H384v640zm1645 158l-557 557l-269-269l90-91l179 179l467-467l90 91z"/>`),
		g.Group(children),
	)
}