package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#000" stroke-linecap="round" d="M4 42H44"/><rect width="12" height="20" x="8" y="22" fill="#2F88FF" stroke="#000" rx="2"/><rect width="20" height="38" x="20" y="4" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M28 32.0083H32"/><path stroke="#fff" stroke-linecap="round" d="M12 32.0083H16"/><path stroke="#fff" stroke-linecap="round" d="M28 23.0083H32"/><path stroke="#fff" stroke-linecap="round" d="M28 14.0083H32"/></g>`),
		g.Group(children),
	)
}