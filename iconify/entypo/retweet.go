package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retweet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 13V8h2L3.5 4L0 8h2v6a2 2 0 0 0 2 2h9.482l-2.638-3H5zm4.156-6L6.518 4H16c1.104 0 2 .897 2 2v6h2l-3.5 4l-3.5-4h2V7H9.156z"/>`),
		g.Group(children),
	)
}