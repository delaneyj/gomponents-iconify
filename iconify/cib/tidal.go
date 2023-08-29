package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tidal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16.016 5.323l-5.339 5.339l-5.339-5.339l-5.339 5.339l5.339 5.339l5.339-5.339l5.339 5.339l-5.339 5.339l5.339 5.339l5.339-5.339l-5.339-5.339l5.339-5.339zm5.375 5.338l5.302-5.307L32 10.661l-5.307 5.307z"/>`),
		g.Group(children),
	)
}