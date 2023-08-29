package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1920h1920v128H0v-128zM128 128H0V0h128v128zM1920 0v128h-128V0h128zM512 128H256V0h256v128zm384 0H640V0h256v128zm384 0h-256V0h256v128zm384 0h-256V0h256v128zM685 813l-90-90l365-365l365 365l-90 90l-211-210v1061H896V603L685 813z"/>`),
		g.Group(children),
	)
}