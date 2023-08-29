package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Credit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M38 15V7H10V15"/><path d="M43 27V15H5V41H43"/><path d="M29 27L24 34L43 34"/><path d="M26 21H22"/><path d="M16 21H12"/></g>`),
		g.Group(children),
	)
}