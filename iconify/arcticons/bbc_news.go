package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BbcNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 18.772h13.138V31.91H5.5zM21.453 5.634H42.5v21.047H21.453zm0 23.862h8.043v12.87h-8.043z"/>`),
		g.Group(children),
	)
}