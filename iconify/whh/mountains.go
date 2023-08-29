package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mountains(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1001 895H24q-40-37-13-62L410 15q16-15 38-15t38 15l178 364l34-108q15-15 37.5-15t38.5 15l239 562q27 25-12 62zM448 128L321 383q0 7-1.5 20t-1 19t3 13.5t10 10T353 447q21-1 37.5-11.5t26-22t26-21T480 383q18 0 34 13.5t28.5 33t26 38.5t32 31.5T640 511z"/>`),
		g.Group(children),
	)
}