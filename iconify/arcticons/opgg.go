package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opgg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.922 29.518a4.272 4.272 0 0 0-4.497-4.271a4.437 4.437 0 0 0-4.038 4.513v3.963A4.272 4.272 0 0 0 17.655 38h0a4.272 4.272 0 0 0 4.267-4.277h-4.267m16.958-4.205a4.272 4.272 0 0 0-4.497-4.271a4.437 4.437 0 0 0-4.038 4.513v3.963A4.272 4.272 0 0 0 30.345 38h0a4.272 4.272 0 0 0 4.268-4.277h-4.268M26.16 22.76V10h4.177a4.285 4.285 0 0 1 0 8.57H26.16m-12.732-.037a4.226 4.226 0 0 0 8.453 0v-4.307a4.226 4.226 0 1 0-8.453 0Z"/>`),
		g.Group(children),
	)
}