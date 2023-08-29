package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 22a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0-2a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm17 13a9 9 0 0 0-9-9H16a9 9 0 1 0 0 18h8a4 4 0 0 0 .883-7.902L28.173 26H32a7 7 0 1 1 0 14v-4a3 3 0 1 0 0-6h-2v12h2a9 9 0 0 0 9-9Zm-25-7a7 7 0 0 0-5.184 11.704l2.7-3.023A3 3 0 0 1 16 30h1.694l3.572-4H16Zm8.11 10.003L22.485 40H24a2 2 0 0 0 .11-3.997ZM20.326 40l5.687-14h-2.066l-11.59 12.98A6.967 6.967 0 0 0 16 40h4.327ZM33 33a1 1 0 0 1-1 1v-2a1 1 0 0 1 1 1Zm-17.095-.996A1 1 0 0 0 15 33v.018l.905-1.014Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}