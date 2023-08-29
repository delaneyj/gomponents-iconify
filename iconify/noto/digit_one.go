package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-digit-onea" gradientUnits="userSpaceOnUse" x1="59.335" y1="100.333" x2="59.335" y2="25.95" gradientTransform="matrix(1 0 0 -1 0 128)"><stop offset="0" stop-color="#757575"/><stop offset=".515" stop-color="#504f4f"/></linearGradient><path d="M73.57 102.07H63.7c-.76 0-1.37-.61-1.37-1.37V43.03c0-.94-.92-1.6-1.81-1.3l-14.98 5.11c-.89.3-1.81-.36-1.81-1.3v-7.76c0-.57.36-1.09.9-1.29L73.1 26.02c.15-.06.31-.08.47-.08c.76 0 1.37.61 1.37 1.37v73.39c0 .75-.61 1.37-1.37 1.37z" fill="url(#ssvg-id-digit-onea)"/>`),
		g.Group(children),
	)
}