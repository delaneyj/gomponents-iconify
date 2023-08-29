package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxSuccessR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" stroke-linecap="round" d="M4 31H15L17 35H31L33 31H44"/><path stroke="#000" stroke-linecap="round" d="M42 36V26"/><path stroke="#000" stroke-linecap="round" d="M6 36V26"/><path stroke="#fff" stroke-linecap="round" d="M17 18.6154L22.6 24L33 14"/></g>`),
		g.Group(children),
	)
}