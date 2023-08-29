package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalTabKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M475 512h1573v128H475l210 211l-90 90l-365-365l365-365l90 90l-210 211zM128 256v640H0V256h128zm1235 941l90-90l365 365l-365 365l-90-90l210-211H0v-128h1573l-210-211zm557 595v-640h128v640h-128z"/>`),
		g.Group(children),
	)
}