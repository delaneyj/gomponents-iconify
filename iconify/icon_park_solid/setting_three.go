package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSettingThree0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="15" fill="#fff"/><path d="M24 39v5m15-21h5M4 23h5M24 4v5m10.606 25.606l3.536 3.536M33.9 12.687l3.535-3.536M9.15 37.435l3.536-3.535M9.858 9.858l3.535 3.535"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSettingThree0)"/>`),
		g.Group(children),
	)
}