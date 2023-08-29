package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudsmith(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M332.5 419.9c0 46.4-37.6 84.1-84 84.1s-84-37.7-84-84.1s37.6-84 84-84s84 37.6 84 84zm-84-243.9c46.4 0 80-37.6 80-84s-33.6-84-80-84s-88 37.6-88 84s-29.6 76-76 76s-84 41.6-84 88s37.6 80 84 80s84-33.6 84-80s33.6-80 80-80z"/>`),
		g.Group(children),
	)
}