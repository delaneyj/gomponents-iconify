package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="26" height="40" x="11" y="4" rx="2"/><circle cx="24" cy="34" r="4"/><rect width="12" height="8" x="18" y="10" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 24H25"/></g>`),
		g.Group(children),
	)
}