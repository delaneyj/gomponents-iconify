package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Navigation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4.805 3.555l15.81 4.865c.862.265.96 1.446.153 1.85l-6.7 3.35a1 1 0 0 0-.448.447l-3.35 6.7c-.403.807-1.584.71-1.85-.153L3.555 4.804a1 1 0 0 1 1.25-1.249Z"/>`),
		g.Group(children),
	)
}