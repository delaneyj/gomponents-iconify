package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoxMultipleTwotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="62" stroke-dashoffset="62" d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/><animate fill="freeze" attributeName="fill-opacity" begin="1.8s" dur="0.15s" values="0;0.3"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M10 6h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M10 10h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M10 14h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path><path stroke-dasharray="34" stroke-dashoffset="34" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.4s" values="34;68"/></path></g>`),
		g.Group(children),
	)
}