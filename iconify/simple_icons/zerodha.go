package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zerodha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.378 5.835A27.267 27.267 0 0 1 24 12.169V0H13.666c2.486 1.343 4.763 3.308 6.712 5.835zM5.48 1.297c-1.914 0-3.755.409-5.48 1.166V24h22.944C22.766 11.44 15 1.297 5.48 1.297z"/>`),
		g.Group(children),
	)
}