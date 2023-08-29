package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.23 37.154a3 3 0 0 0 5.54 0l10-24a3 3 0 0 0-5.54-2.308L24 28.2l-7.23-17.354a3 3 0 0 0-5.54 2.308l10 24ZM24 37a1 1 0 0 1-.923-.615l-10-24a1 1 0 1 1 1.846-.77l8.154 19.57a1 1 0 0 0 1.846 0l8.154-19.57a1 1 0 0 1 1.846.77l-10 24A1 1 0 0 1 24 37Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}