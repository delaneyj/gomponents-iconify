package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDiskLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M217.9 81.42L174.58 38.1a13.9 13.9 0 0 0-9.89-4.1H48a14 14 0 0 0-14 14v160a14 14 0 0 0 14 14h160a14 14 0 0 0 14-14V91.31a13.9 13.9 0 0 0-4.1-9.89ZM170 210H86v-58a2 2 0 0 1 2-2h80a2 2 0 0 1 2 2Zm40-2a2 2 0 0 1-2 2h-26v-58a14 14 0 0 0-14-14H88a14 14 0 0 0-14 14v58H48a2 2 0 0 1-2-2V48a2 2 0 0 1 2-2h116.69a2 2 0 0 1 1.41.58l43.32 43.32a2 2 0 0 1 .58 1.41ZM158 72a6 6 0 0 1-6 6H96a6 6 0 0 1 0-12h56a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}