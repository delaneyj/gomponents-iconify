package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h1920V0H0v128zm128 1792H0v128h128v-128zm1792 128v-128h-128v128h128zM512 1920H256v128h256v-128zm384 0H640v128h256v-128zm384 0h-256v128h256v-128zm384 0h-256v128h256v-128zm-979-685l-90 90l365 365l365-365l-90-90l-211 210V384H896v1061l-211-210z"/>`),
		g.Group(children),
	)
}