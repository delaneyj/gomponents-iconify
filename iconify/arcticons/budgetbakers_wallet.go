package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BudgetbakersWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.451 29.914v5.893a3.946 3.946 0 0 1-3.946 3.946H8.446A3.946 3.946 0 0 1 4.5 35.808V12.193a3.946 3.946 0 0 1 3.946-3.947h30.059a3.946 3.946 0 0 1 3.946 3.947v6.098"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.294 29.925H31.925a5.822 5.822 0 0 1-5.822-5.822h0a5.822 5.822 0 0 1 5.822-5.823h10.37c.665 0 1.205.54 1.205 1.206v9.234c0 .665-.54 1.205-1.206 1.205Z"/><circle cx="32.458" cy="24.171" r="2.705" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}