package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundwaveBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 4v16m4-13v2m0 8v-4m-8 4v-2m0-8v4m12 0v2M4 11v2"/>`),
		g.Group(children),
	)
}