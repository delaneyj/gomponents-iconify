package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartflasher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.66 4.5V24h-6.33l11.01 19.5V24h6.33L21.66 4.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.34 36.55a12.76 12.76 0 0 0-.83-25.22m-3.85.12a12.76 12.76 0 0 0 .83 25.22"/>`),
		g.Group(children),
	)
}