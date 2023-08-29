package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-dark-skin-tonea" x1="64" x2="64" y1="4.37" y2="123.52" gradientUnits="userSpaceOnUse"><stop stop-color="#70534A" offset="0"/><stop stop-color="#6D5047" offset=".467"/><stop stop-color="#63463D" offset=".842"/><stop stop-color="#5C4037" offset="1"/></linearGradient><path fill="url(#ssvg-id-dark-skin-tonea)" d="M4 4h120v120H4z"/>`),
		g.Group(children),
	)
}