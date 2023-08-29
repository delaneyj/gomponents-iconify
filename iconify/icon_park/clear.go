package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#000" stroke-linecap="round" d="M20 5.91406H28V13.9141H43V21.9141H5V13.9141H20V5.91406Z" clip-rule="evenodd"/><path fill="#2F88FF" stroke="#000" d="M8 40H40V22H8V40Z"/><path stroke="#fff" stroke-linecap="round" d="M16 39.8976V33.9141"/><path stroke="#fff" stroke-linecap="round" d="M24 39.8977V33.8977"/><path stroke="#fff" stroke-linecap="round" d="M32 39.8976V33.9141"/><path stroke="#000" stroke-linecap="round" d="M12 40H36"/></g>`),
		g.Group(children),
	)
}