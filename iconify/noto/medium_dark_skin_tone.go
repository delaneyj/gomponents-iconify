package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediumDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-medium-dark-skin-tonea" x1="64" x2="64" y1="4.333" y2="123.89" gradientUnits="userSpaceOnUse"><stop stop-color="#A56C43" offset="0"/><stop stop-color="#A26941" offset=".427"/><stop stop-color="#97603D" offset=".784"/><stop stop-color="#8D5738" offset="1"/></linearGradient><path fill="url(#ssvg-id-medium-dark-skin-tonea)" d="M4 4h120v120H4z"/><linearGradient id="ssvg-id-medium-dark-skin-toneb" x1="106" x2="106" y1="305.27" y2="305.27" gradientUnits="userSpaceOnUse"><stop stop-color="#8D5738" offset="0"/><stop stop-color="#97603D" offset=".216"/><stop stop-color="#A26941" offset=".573"/><stop stop-color="#A56C43" offset="1"/></linearGradient>`),
		g.Group(children),
	)
}