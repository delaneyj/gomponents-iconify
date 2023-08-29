package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-light-skin-tonea" x1="64" x2="64" y1="4.5" y2="123.82" gradientUnits="userSpaceOnUse"><stop stop-color="#F9DDBD" offset="0"/><stop stop-color="#FADCBA" offset=".37"/><stop stop-color="#FCD8AF" offset=".684"/><stop stop-color="#FFD39E" offset=".976"/><stop stop-color="#FFD29C" offset="1"/></linearGradient><path fill="url(#ssvg-id-light-skin-tonea)" d="M4 4h120v120H4z"/>`),
		g.Group(children),
	)
}