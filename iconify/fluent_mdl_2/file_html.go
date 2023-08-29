package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHTML(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1920h256v128H128V0h1115l512 512h37v768h-128V640h-512V128H256v1792zM1573 512l-293-293v293h293zM877 1517l-211 211l211 211l-90 90l-301-301l301-301l90 90zm339 19q40 0 75 15t61 41t41 61t15 75q0 40-15 75t-41 61t-61 41t-75 15q-40 0-75-15t-61-41t-41-61t-15-75q0-40 15-75t41-61t61-41t75-15zm730 192l-301 301l-90-90l211-211l-211-211l90-90l301 301z"/>`),
		g.Group(children),
	)
}