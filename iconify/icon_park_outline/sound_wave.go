package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 14v-2a6 6 0 0 1 6-6h24a6 6 0 0 1 6 6v2m-10 4v12m8-10v8M24 15v18m-8-15v12M8 20v8m-2 6v2a6 6 0 0 0 6 6h24a6 6 0 0 0 6-6v-2"/>`),
		g.Group(children),
	)
}