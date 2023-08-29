package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M5 6C5 5.44772 5.44772 5 6 5L16 5C16.5523 5 17 5.44772 17 6L17 42C17 42.5523 16.5523 43 16 43H6C5.44772 43 5 42.5523 5 42L5 6Z"/><circle cx="11" cy="35" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M11 10.9999V24.9999"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M23.9775 9.75562L34.0372 7.24071C34.5701 7.10747 35.1107 7.42907 35.248 7.96097L42.7482 37.0244C42.887 37.562 42.5611 38.1097 42.0225 38.2444L31.9628 40.7593C31.4299 40.8925 30.8893 40.5709 30.752 40.039L23.2518 10.9756C23.113 10.438 23.4389 9.89029 23.9775 9.75562Z"/><circle cx="35" cy="32" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 16L33.5 26"/></g>`),
		g.Group(children),
	)
}