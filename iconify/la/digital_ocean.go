package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalOcean(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.831 3 3 8.832 3 16h5a8 8 0 1 1 8 8v5c7.169 0 13-5.832 13-13S23.169 3 16 3zm0 21v-6h-6v6h6zm-6 0H6v4h4v-4zm-4 0v-3H3v3h3z"/>`),
		g.Group(children),
	)
}