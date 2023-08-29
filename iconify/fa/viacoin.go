package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viacoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1536 0l-192 448h192v192h-274l-55 128h329v192h-411l-357 832l-357-832H0V768h329l-55-128H0V448h192L0 0h256l323 768h378L1280 0h256zM768 1216l108-256H660z"/>`),
		g.Group(children),
	)
}