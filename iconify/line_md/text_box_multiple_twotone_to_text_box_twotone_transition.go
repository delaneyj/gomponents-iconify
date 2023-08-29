package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoxMultipleTwotoneToTextBoxTwotoneTransition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity=".3" d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17zM10 6h8M10 10h8M10 14h5"><animateMotion fill="freeze" begin="0.4s" calcMode="linear" dur="0.4s" path="M0 0l-2 2"/></path><path stroke-dasharray="34" stroke-dashoffset="68" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="68;34"/></path></g>`),
		g.Group(children),
	)
}