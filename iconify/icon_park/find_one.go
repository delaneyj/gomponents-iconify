package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FindOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M23.8 14C29.4333 14 34 18.4772 34 24C34 29.5228 29.4333 34 23.8 34C21.1876 34 18.8046 33.0372 17 31.4537"/><path d="M24 4C35.0457 4 44 12.9543 44 24C44 35.0457 35.0457 44 24 44C18.5856 44 13.6737 41.8485 10.0726 38.3537C6.32733 34.719 4 29.6313 4 24C4 18.6326 6.11433 13.759 9.55556 10.1667"/><path d="M9.55566 10.167L24.0001 24.0003L10.0727 38.3539"/></g>`),
		g.Group(children),
	)
}