package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediumLightSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-medium-light-skin-tonea" x1="64" x2="64" y1="5" y2="123.6" gradientUnits="userSpaceOnUse"><stop stop-color="#E0BB95" offset="0"/><stop stop-color="#DEB892" offset=".411"/><stop stop-color="#D6B088" offset=".743"/><stop stop-color="#CCA47A" offset="1"/></linearGradient><path fill="url(#ssvg-id-medium-light-skin-tonea)" d="M4 4h120v120H4z"/>`),
		g.Group(children),
	)
}