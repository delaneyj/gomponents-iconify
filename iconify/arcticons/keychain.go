package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keychain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.35 4.5a2 2 0 0 0-1.95 2v35.1a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95Zm7.19 5.21a3.75 3.75 0 0 1 3.75 3.75h0a3.75 3.75 0 1 1-3.75-3.75ZM13.8 24h20.4m-20.4 6.32h20.4m-20.4 6.32h20.4"/>`),
		g.Group(children),
	)
}