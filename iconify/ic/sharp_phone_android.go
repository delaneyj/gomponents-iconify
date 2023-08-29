package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPhoneAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 1H5v22h14V1zm-5 20h-4v-1h4v1zm3-3H7V4h10v14z"/>`),
		g.Group(children),
	)
}