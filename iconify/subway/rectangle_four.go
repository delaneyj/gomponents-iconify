package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M47.5 280.3H1v93.1h46.5v-93.1zM1 94.1h46.5V47.5H94V1H1v93.1zm46.5 325.8H1V513h93.1v-46.5H47.5v-46.6zM373.4 1h-93.1v46.5h93.1V1zM47.5 140.6H1v93.1h46.5v-93.1zM140.6 513h93.1v-46.5h-93.1V513zM419.9 1v46.5h46.5V94H513V1h-93.1zM140.6 47.5h93.1V1h-93.1v46.5zm325.9 186.2H513v-93.1h-46.5v93.1zM280.3 513H513V280.3H280.3V513z"/>`),
		g.Group(children),
	)
}