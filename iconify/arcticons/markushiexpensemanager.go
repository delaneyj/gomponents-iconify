package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markushiexpensemanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.961 13.065h7.518v21.87H6.961zm13.832 5.467h7.518v16.402h-7.518zm14.201 8.885h7.518v7.518h-7.518z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 34.924l14.294-11.49L31.27 27.93l12.23-9.427"/>`),
		g.Group(children),
	)
}