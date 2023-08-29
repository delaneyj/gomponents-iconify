package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NordicNoodle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.75 34.25v-20.5L24 3.5L6.25 13.75v20.5L24 44.5l17.75-10.25zm-8.85-.86L14.61 15.1m.28 18.01l18.55-18.55"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.61 16.73l-3.92 3.92c-1.2 1.2-3.14 1.2-4.34 0h0c-1.2-1.2-1.2-3.14 0-4.34l3.92-3.92"/><circle cx="23.76" cy="13.72" r=".75" fill="currentColor"/><circle cx="34.28" cy="23.76" r=".75" fill="currentColor"/><circle cx="24.24" cy="34.28" r=".75" fill="currentColor"/><circle cx="13.72" cy="24.24" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}