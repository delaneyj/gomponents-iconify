package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingConfig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M41.5 10H35.5"/><path d="M27.5 6V14"/><path d="M27.5 10L5.5 10"/><path d="M13.5 24H5.5"/><path d="M21.5 20V28"/><path d="M43.5 24H21.5"/><path d="M41.5 38H35.5"/><path d="M27.5 34V42"/><path d="M27.5 38H5.5"/></g>`),
		g.Group(children),
	)
}