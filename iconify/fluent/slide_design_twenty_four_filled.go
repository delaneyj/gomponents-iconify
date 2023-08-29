package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideDesignTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M16 3a1 1 0 0 1 1 1v6h2a1 1 0 1 1 0 2h-2v5.05a2.512 2.512 0 0 0-2 .45V4a1 1 0 0 1 1-1zm.5 15a1.5 1.5 0 1 1-.001 3a1.5 1.5 0 0 1 .001-3zM4 5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1c0 1.77-.366 4.436-1.507 6.7C10.335 13.997 8.299 16 5 16a1 1 0 1 1 0-2c2.301 0 3.764-1.33 4.707-3.2c.771-1.53 1.138-3.338 1.252-4.8H5a1 1 0 0 1-1-1zm8.5 16a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm9.5-1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}