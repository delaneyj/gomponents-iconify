package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnOffBluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M20.6719 11.7778V4L35.2563 15.1111L28.0771 19.5556M32.4996 35.9744L20.6751 44V29.6275L32.4996 35.9744Z"/><path d="M4 12.5L44 35.5"/><path stroke-linejoin="round" d="M7.44434 33.9999L20.6751 29.6274"/></g>`),
		g.Group(children),
	)
}