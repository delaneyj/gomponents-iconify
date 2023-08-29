package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 640v128H0V640l960-480l960 480zM256 896v512H128V896h128zm256 0v512H384V896h128zm256 0v512H640V896h128zm256 0v512H896V896h128zm256 0v512h-128V896h128zm256 0v512h-128V896h128zm256 512h-128V896h128v512zm0 128l128 384H0l128-384h1664z"/>`),
		g.Group(children),
	)
}