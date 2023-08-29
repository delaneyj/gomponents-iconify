package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircularSaw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16v72H88l48 48L16 256h72v168l48-48l120 120v-72h168l-48-48l120-120h-72V88l-48 48L256 16zm0 120c66.274 0 120 53.726 120 120s-53.726 120-120 120s-120-53.726-120-120s53.726-120 120-120zm1.406 72.03A48 48 0 0 0 208 256a48 48 0 0 0 96 0a48 48 0 0 0-46.594-47.97z"/>`),
		g.Group(children),
	)
}