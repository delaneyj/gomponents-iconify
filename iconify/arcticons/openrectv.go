package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openrectv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.201 23.134L21.2 17.357a1 1 0 0 0-1.5.866v11.554a1 1 0 0 0 1.5.866l10-5.777a1 1 0 0 0 .001-1.732Z"/><circle cx="24" cy="24" r="14.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5a21.477 21.477 0 1 1-6.117.883L24 9.5Z"/>`),
		g.Group(children),
	)
}