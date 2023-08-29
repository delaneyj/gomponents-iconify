package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridNineBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 44H40a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h176a20 20 0 0 0 20-20V64a20 20 0 0 0-20-20Zm-108 96v-24h40v24Zm40 24v24h-40v-24ZM44 116h40v24H44Zm64-24V68h40v24Zm64 24h40v24h-40Zm40-24h-40V68h40ZM84 68v24H44V68Zm-40 96h40v24H44Zm128 24v-24h40v24Z"/>`),
		g.Group(children),
	)
}