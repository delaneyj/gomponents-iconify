package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppVault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 11h-33a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2V13a2 2 0 0 0-2-2ZM39 10V6.5a1 1 0 0 0-1-1H10a1 1 0 0 0-1 1V10a1 1 0 0 0 1 1h28a1 1 0 0 0 1-1ZM9 38v3.5a1 1 0 0 0 1 1h28a1 1 0 0 0 1-1V38a1 1 0 0 0-1-1H10a1 1 0 0 0-1 1Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 15h7v7h-7z"/>`),
		g.Group(children),
	)
}