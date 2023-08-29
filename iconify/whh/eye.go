package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 768q-122 0-231-50T97.5 580T0 384q23-108 97.5-196T281 50T512 0t231 50t183.5 138t97.5 196q-23 108-97.5 196T743 718t-231 50zM266 180q-78 35-131 88T64 384q18 64 71.5 117T266 589q-74-89-74-205t74-204zm246 76q-53 0-90.5 37.5T384 384t37.5 90.5T512 512t90.5-37.5T640 384t-37.5-90.5T512 256zm246-76q74 88 74 204t-74 205q78-35 131-88t71-117q-18-63-71-116.5T758 180z"/>`),
		g.Group(children),
	)
}