package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeckboardPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.509 25.448l-8.536-4.928l-8.482 4.898l8.536 4.928l8.482-4.898z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.491 25.418v9.825l8.536 4.928l8.482-4.897v-9.826M4.5 19.072v9.825l8.536 4.928l2.455-1.417m17.018-.031l2.509 1.448l8.482-4.897v-9.825m0 0l-8.536-4.928l-8.482 4.897L35.018 24l8.482-4.897zm-21.982 0l-8.536-4.928L4.5 19.072L13.036 24l8.482-4.897zm10.991-6.346l-8.536-4.928l-8.482 4.897l8.536 4.928l8.482-4.897zm0 0v2.835m-17.018-2.866v2.866m10.991 3.48v2.896m-4.964-2.865v2.865"/>`),
		g.Group(children),
	)
}