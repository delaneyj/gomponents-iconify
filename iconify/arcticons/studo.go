package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Studo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="32.594" height="32.594" x="7.703" y="7.703" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6" transform="rotate(45 24 24)"/><circle cx="24" cy="24" r="4.346" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.927 20.927l-8.451-8.451"/>`),
		g.Group(children),
	)
}