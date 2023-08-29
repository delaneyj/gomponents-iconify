package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M22 13h-2v2h2v-2Zm2 0h2v2h-2v-2Zm6 0h-2v2h2v-2Zm-10 4h2v2h-2v-2Zm6 0h-2v2h2v-2Zm2 0h2v2h-2v-2Zm-6 4h-2v2h2v-2Zm2 0h2v2h-2v-2Zm6 0h-2v2h2v-2Zm-10 4h2v2h-2v-2Zm6 0h-2v2h2v-2Zm-6 4h2v2h-2v-2Zm6 0h-2v2h2v-2Zm-6 4h2v2h-2v-2Zm6 0h-2v2h2v-2Zm-6 4h2v2h-2v-2Zm6 0h-2v2h2v-2Zm5-7h3v-2h-3v2Zm3 4h-3v-2h3v2Zm-3 4h3v-2h-3v2Z"/><path fill-rule="evenodd" d="m17 4l16 6v14h4a1 1 0 0 1 1 1v17h1a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h1V21a1 1 0 0 1 1-1h2v-7h2v7h2V4Zm2 2.886l12 4.5V24h-3a1 1 0 0 0-1 1v17h-8V6.886ZM12 22v20h5V22h-5Zm24 20h-2v-2h-3v2h-2V26h7v16Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}