package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartTransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="17.416" height="30.217" x="21.156" y="13.283" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.597"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.156 16.677h17.416M21.156 40.106h17.416m-17.416-5.389h-10.13a1.597 1.597 0 0 1-1.598-1.597V6.098A1.597 1.597 0 0 1 11.025 4.5h14.221a1.597 1.597 0 0 1 1.598 1.597v7.186M9.428 7.894h17.416M9.428 31.323h11.728m6.047-1.276L15.109 17.953m12.094 8.124v3.97h-3.97"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.109 21.923v-3.97h3.97"/>`),
		g.Group(children),
	)
}