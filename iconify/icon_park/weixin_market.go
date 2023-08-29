package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinMarket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M38 14H10C8.89543 14 8 14.8954 8 16V42C8 43.1046 8.89543 44 10 44H38C39.1046 44 40 43.1046 40 42V16C40 14.8954 39.1046 14 38 14Z"/><path d="M17 18V11C17 7.13401 20.134 4 24 4C27.866 4 31 7.13401 31 11V18"/></g>`),
		g.Group(children),
	)
}