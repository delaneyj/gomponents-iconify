package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="34.132" x="5.5" y="6.934" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.995"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 26.6h37m-37 4.483h37"/>`),
		g.Group(children),
	)
}