package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 8.5v31"/><circle cx="19.5" cy="24" r="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.748 27.11a4.5 4.5 0 1 0 0-6.22"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.248 27.11a4.5 4.5 0 1 0 0-6.221"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.747 27.11a4.5 4.5 0 1 0 .071-6.292"/>`),
		g.Group(children),
	)
}