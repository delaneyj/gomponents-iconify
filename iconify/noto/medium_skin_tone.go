package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediumSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-medium-skin-tonea" x1="64" x2="64" y1="3.667" y2="123.67" gradientUnits="userSpaceOnUse"><stop stop-color="#BA8D68" offset="0"/><stop stop-color="#B78A67" offset=".449"/><stop stop-color="#AD8264" offset=".809"/><stop stop-color="#A47B62" offset="1"/></linearGradient><path fill="url(#ssvg-id-medium-skin-tonea)" d="M4 4h120v120H4z"/>`),
		g.Group(children),
	)
}