package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.36 21.44L13.174 3.996l2.101-3.6l-.686-.4l-1.872 3.213l-1.872-3.213l-.686.4l2.101 3.6L2.081 21.44H.001V24h25.446v-2.56zm-12.286-5.06a14.947 14.947 0 0 0 1.621-3.949l.023-.105c.966 3.835 3.282 7.017 6.401 9.085l.057.035H6.269a15.495 15.495 0 0 0 4.771-4.993l.039-.071z"/>`),
		g.Group(children),
	)
}