package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Padding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0h128v128H0V0zm1920 0v128h-128V0h128zM256 0h256v128H256V0zm640 128H640V0h256v128zm384 0h-256V0h256v128zM1408 0h256v128h-256V0zM0 1920h128v128H0v-128zm1792 0h128v128h-128v-128zm-1536 0h256v128H256v-128zm384 0h256v128H640v-128zm384 0h256v128h-256v-128zm384 0h256v128h-256v-128zM1024 603v842l211-210l90 90l-365 365l-365-365l90-90l211 210V603L685 813l-90-90l365-365l365 365l-90 90l-211-210z"/>`),
		g.Group(children),
	)
}