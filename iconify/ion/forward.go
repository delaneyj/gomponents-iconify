package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M288 298.1v92.3L448 256 288 112v80C100.8 192 64 400 64 400c53-93 122.4-101.9 224-101.9z" fill="currentColor"/>`),
		g.Group(children),
	)
}