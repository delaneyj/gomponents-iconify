package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boxing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" stroke-miterlimit="2" d="M38 36V42C38 43.1 37.11 44 36 44H17C15.9 44 15 43.11 15 42V36"/><path fill="#2F88FF" stroke-linejoin="round" stroke-miterlimit="2" d="M11 15V10C11 6.68629 13.6863 4 17 4H28H38C41.3137 4 44 6.68629 44 10V30C44 33.3137 41.3198 36 38.006 36C30.9668 36 19.6598 36 15 36C8 36 4 28 4 23C4 18 4 15 4 15H11Z"/><line x1="11" x2="11" y1="14" y2="22"/></g>`),
		g.Group(children),
	)
}