package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageService(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 20h-2v2h2v6H4v-6h2v-2H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2Z"/><circle cx="7" cy="25" r="1" fill="currentColor"/><path fill="currentColor" d="M13 4h2v8h-2zM9 4h2v8H9zm12 8h-2a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h2a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zm-2-6v4h2V6zm2 8h2v8h-2zM9 14h2v8H9zm8 8h-2a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h2a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zm-2-6v4h2v-4z"/>`),
		g.Group(children),
	)
}