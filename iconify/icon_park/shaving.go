package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shaving(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="20" height="27" x="14" y="17" fill="#2F88FF" rx="2"/><rect width="10" height="5" x="19" y="12"/><path d="M19 12C19 10 19 10 19.0001 9C19.0001 8 19 4 27 4C35 4 36 4 36 4V9H29V12"/></g>`),
		g.Group(children),
	)
}