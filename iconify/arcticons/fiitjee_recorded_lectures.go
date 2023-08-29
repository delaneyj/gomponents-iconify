package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiitjeeRecordedLectures(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.787 5.676h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM39.273 15.97v-5a2 2 0 0 0-2-2h-7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.81 24.2l-14.97 8.643V15.556ZM9 36.706v2.085a.834.834 0 0 0 .834.835h2.92"/>`),
		g.Group(children),
	)
}