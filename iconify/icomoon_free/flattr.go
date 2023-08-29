package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flattr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.743 0C1.941 0 0 2.19 0 6.279v8.579l3.725-3.729V6.771c0-1.694.449-2.772 1.955-3.014c.526-.103 1.621-.067 2.317-.067v2.587a.247.247 0 0 0 .245.269c.063 0 .123-.033.184-.093L14.881 0L5.742-.001zm6.532 4.871v4.358c0 1.694-.449 2.772-1.955 3.014c-.526.103-1.621.067-2.317.067V9.723a.414.414 0 0 0-.009-.087a.246.246 0 0 0-.236-.182c-.064 0-.123.033-.184.093L1.119 16l9.139.001c3.802 0 5.743-2.19 5.743-6.279V1.143l-3.725 3.729z"/>`),
		g.Group(children),
	)
}