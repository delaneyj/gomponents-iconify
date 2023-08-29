package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchContrast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 10h10c1.886 0 2.828 0 3.414.586C44 11.172 44 12.114 44 14v20c0 1.886 0 2.828-.586 3.414C42.828 38 41.886 38 40 38H30M18 10H8c-1.886 0-2.828 0-3.414.586C4 11.172 4 12.114 4 14v20c0 1.886 0 2.828.586 3.414C5.172 38 6.114 38 8 38h10m6-32v36"/>`),
		g.Group(children),
	)
}