package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoFasAuth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.56 42.945l9.925-8.088a9.028 9.028 0 0 0 3.326-7V9.708a2.657 2.657 0 0 0-1.975-2.568l-7.15-1.897a22.166 22.166 0 0 0-11.372 0l-7.15 1.897a2.657 2.657 0 0 0-1.975 2.568v18.15a9.028 9.028 0 0 0 3.325 7l9.926 8.088a2.468 2.468 0 0 0 3.12 0ZM15.049 14.797h17.902m-17.902 5.979H24m-8.951 5.98h2.082"/>`),
		g.Group(children),
	)
}