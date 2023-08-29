package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gratipay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M248 8C111.1 8 0 119.1 0 256s111.1 248 248 248s248-111.1 248-248S384.9 8 248 8zm114.6 226.4l-113 152.7l-112.7-152.7c-8.7-11.9-19.1-50.4 13.6-72c28.1-18.1 54.6-4.2 68.5 11.9c15.9 17.9 46.6 16.9 61.7 0c13.9-16.1 40.4-30 68.1-11.9c32.9 21.6 22.6 60 13.8 72z"/>`),
		g.Group(children),
	)
}