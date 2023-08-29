package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneHandLiquidSoap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22.012 14.74a3.504 3.504 0 0 1-7.008 0c0-2.628 3.5-7.009 3.5-7.009s3.508 4.381 3.508 7.009ZM9.998 9.233H3.99a2.002 2.002 0 0 0-2.002 2.002v10.013c0 1.106.896 2.002 2.002 2.002h6.008A2.002 2.002 0 0 0 12 21.248V11.235a2.002 2.002 0 0 0-2.002-2.002ZM4.766 6.23h4.456a.776.776 0 0 1 .778.775v2.228H3.99V7.005a.776.776 0 0 1 .776-.775v0ZM14 2.752l-.447-.895A2 2 0 0 0 11.764.75H2.989m4.005 13.489v4.005m-2.002-2.002h4.004M6.994.75v5.48"/>`),
		g.Group(children),
	)
}