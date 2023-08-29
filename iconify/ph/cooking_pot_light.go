package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookingPotLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M90 48V16a6 6 0 0 1 12 0v32a6 6 0 0 1-12 0Zm38 6a6 6 0 0 0 6-6V16a6 6 0 0 0-12 0v32a6 6 0 0 0 6 6Zm32 0a6 6 0 0 0 6-6V16a6 6 0 0 0-12 0v32a6 6 0 0 0 6 6Zm91.6 46.8L222 123v61a30 30 0 0 1-30 30H64a30 30 0 0 1-30-30v-61L4.4 100.8a6 6 0 0 1 7.2-9.6L34 108V88a14 14 0 0 1 14-14h160a14 14 0 0 1 14 14v20l22.4-16.8a6 6 0 0 1 7.2 9.6ZM210 88a2 2 0 0 0-2-2H48a2 2 0 0 0-2 2v96a18 18 0 0 0 18 18h128a18 18 0 0 0 18-18Z"/>`),
		g.Group(children),
	)
}