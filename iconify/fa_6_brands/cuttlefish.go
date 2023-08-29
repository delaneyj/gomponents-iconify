package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cuttlefish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M344 305.5c-17.5 31.6-57.4 54.5-96 54.5c-56.6 0-104-47.4-104-104s47.4-104 104-104c38.6 0 78.5 22.9 96 54.5c13.7-50.9 41.7-93.3 87-117.8C385.7 39.1 320.5 8 248 8C111 8 0 119 0 256s111 248 248 248c72.5 0 137.7-31.1 183-80.7c-45.3-24.5-73.3-66.9-87-117.8z"/>`),
		g.Group(children),
	)
}