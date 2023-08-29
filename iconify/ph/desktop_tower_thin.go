package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopTowerThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M20 96v72a12 12 0 0 0 12 12h80a4 4 0 0 1 0 8H92v24h20a4 4 0 0 1 0 8H64a4 4 0 0 1 0-8h20v-24H32a20 20 0 0 1-20-20V96a20 20 0 0 1 20-20h80a4 4 0 0 1 0 8H32a12 12 0 0 0-12 12Zm188-28h-32a4 4 0 0 0 0 8h32a4 4 0 0 0 0-8Zm0 32h-32a4 4 0 0 0 0 8h32a4 4 0 0 0 0-8Zm36-52v160a12 12 0 0 1-12 12h-80a12 12 0 0 1-12-12V48a12 12 0 0 1 12-12h80a12 12 0 0 1 12 12Zm-8 0a4 4 0 0 0-4-4h-80a4 4 0 0 0-4 4v160a4 4 0 0 0 4 4h80a4 4 0 0 0 4-4Zm-44 124a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}