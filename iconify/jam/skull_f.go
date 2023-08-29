package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkullF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 17.858a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-.865a7.499 7.499 0 0 1-4-6.635v-2a7.5 7.5 0 0 1 7.5-7.5h5a7.5 7.5 0 0 1 7.5 7.5v2a7.499 7.499 0 0 1-4 6.635v.865zm-10-6a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm8 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm-4 4a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}