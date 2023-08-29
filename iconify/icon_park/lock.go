package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="22" x="6" y="22" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#000" stroke-linecap="round" d="M14 22V14C14 8.47715 18.4772 4 24 4C29.5228 4 34 8.47715 34 14V22"/><path stroke="#fff" stroke-linecap="round" d="M24 30V36"/></g>`),
		g.Group(children),
	)
}