package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookingPotThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 48V16a4 4 0 0 1 8 0v32a4 4 0 0 1-8 0Zm36 4a4 4 0 0 0 4-4V16a4 4 0 0 0-8 0v32a4 4 0 0 0 4 4Zm32 0a4 4 0 0 0 4-4V16a4 4 0 0 0-8 0v32a4 4 0 0 0 4 4Zm90.4 47.2L220 122v62a28 28 0 0 1-28 28H64a28 28 0 0 1-28-28v-62L5.6 99.2a4 4 0 0 1 4.8-6.4L36 112V88a12 12 0 0 1 12-12h160a12 12 0 0 1 12 12v24l25.6-19.2a4 4 0 1 1 4.8 6.4ZM212 88a4 4 0 0 0-4-4H48a4 4 0 0 0-4 4v96a20 20 0 0 0 20 20h128a20 20 0 0 0 20-20Z"/>`),
		g.Group(children),
	)
}