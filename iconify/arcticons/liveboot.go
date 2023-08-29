package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liveboot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.884h39m-39 34.232h39m-39-4.279h39M4.5 24h39m-39 4.279h14.838M4.5 32.558h14.838M4.5 15.442h25.365M4.5 11.163h25.365M4.5 19.721h25.365"/>`),
		g.Group(children),
	)
}