package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fengoffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m979 767l-282 89l193-445l127 269q13 27 2 52.5T979 767zM411 134L680 7q27-13 52.5-2T767 45l89 282zM134 613L7 344q-13-27-2-52.5T45 257l282-89zm479 277l-269 127q-27 13-52.5 2T257 979l-89-282z"/>`),
		g.Group(children),
	)
}