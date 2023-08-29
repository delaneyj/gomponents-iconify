package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M97 289q-40 0-68.5-28.5T0 192.5T28.5 125T97 97t68 28t28 67.5t-28 68T97 289zm233 0q-40 0-68.5-28.5t-28.5-68t28.5-67.5T330 97t68.5 28t28.5 67.5t-28.5 68T330 289z"/>`),
		g.Group(children),
	)
}