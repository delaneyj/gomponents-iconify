package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Face(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 211q11 0 19 7.5t8 18.5t-8 19t-19 8t-18.5-8t-7.5-19t7.5-18.5T149 211zm128 0q11 0 19 7.5t8 18.5t-8 19t-19 8t-18.5-8t-7.5-19t7.5-18.5T277 211zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zm0 384q70.5 0 120.5-50t50-121q0-24-7-48q-24 5-48 5q-53 0-99-24t-75-66q-33 80-111 115q-1 10-1 18q0 71 50 121t120.5 50z"/>`),
		g.Group(children),
	)
}