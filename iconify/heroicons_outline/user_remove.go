package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-4 7a6 6 0 0 0-6 6v1h12v-1a6 6 0 0 0-6-6Zm12-2h-6"/>`),
		g.Group(children),
	)
}