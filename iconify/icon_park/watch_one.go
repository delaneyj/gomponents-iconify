package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linecap="round" d="M19 14V6C19 4.89543 19.8954 4 21 4H27C28.1046 4 29 4.89543 29 6V14"/><path stroke="#000" stroke-linecap="round" d="M29 34V42C29 43.1046 28.1046 44 27 44H21C19.8954 44 19 43.1046 19 42V34"/><circle cx="24" cy="24" r="11" fill="#2F88FF" stroke="#000"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M21 25L25 25"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M25 21V25"/></g>`),
		g.Group(children),
	)
}