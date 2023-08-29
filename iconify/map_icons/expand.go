package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M2 15.758V2h14.299l5.262 4h-8.769L22 15.758L16.299 21L7 12.251v8.769zm46 17.994V48H33.701l-5.262-4h8.769L28 33.997l5.701-5.364L43 37.259V28.49z"/>`),
		g.Group(children),
	)
}