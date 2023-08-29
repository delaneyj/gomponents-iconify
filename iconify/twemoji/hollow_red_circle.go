package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HollowRedCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#DD2E44" d="M18 0C8.059 0 0 8.059 0 18s8.059 18 18 18s18-8.059 18-18S27.941 0 18 0zm0 30c-6.627 0-12-5.373-12-12S11.373 6 18 6s12 5.373 12 12s-5.373 12-12 12z"/>`),
		g.Group(children),
	)
}