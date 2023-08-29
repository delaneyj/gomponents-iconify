package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.485 29.943v-12.11m0 7.626A4.5 4.5 0 0 1 24 29.943h0a4.5 4.5 0 0 1-4.485-4.486v-2.915A4.5 4.5 0 0 1 24 18.058h0a4.5 4.5 0 0 1 4.485 4.485m-13.977-8.748L4.5 23.804l10.402 10.402m18.59-20.412L43.5 23.803L33.098 34.205"/>`),
		g.Group(children),
	)
}