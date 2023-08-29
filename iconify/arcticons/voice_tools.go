package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36 24.12l1.62-1.62l5.88 5.88M33 21.12l4.62-4.62l5.88 5.88m-39 3.144l5.318 5.318L13.16 27.5M4.5 19.524l5.318 5.318L15.16 19.4m-1.165 4.421a10.2 10.2 0 0 0 20.01 0M24 7.425h0a6.468 6.468 0 0 1 6.468 6.469V21.9A6.468 6.468 0 0 1 24 28.37h0a6.468 6.468 0 0 1-6.468-6.468v-8.007A6.468 6.468 0 0 1 24 7.425Zm2.499 24.352v8.254c0 .297-.238.54-.535.544h-3.927a.544.544 0 0 1-.536-.544v-8.254"/>`),
		g.Group(children),
	)
}