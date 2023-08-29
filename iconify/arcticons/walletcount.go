package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walletcount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 17V7.5a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2V31"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.486 17H29a.757.757 0 0 0-.757.757v12.486A.757.757 0 0 0 29 31h15.486a.757.757 0 0 0 .757-.757V17.757a.757.757 0 0 0-.757-.757Z"/><circle cx="34.743" cy="24" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}