package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 0q136 0 256 91L256 405L0 90Q119 0 256 0zm21 277V149h-42v128h42zm-42-170h42V64h-42v43z"/>`),
		g.Group(children),
	)
}