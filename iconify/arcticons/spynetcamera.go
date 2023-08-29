package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spynetcamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.95" cy="23.95" r="7.88" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="bevel"/><circle cx="23.95" cy="23.95" r="11.37" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="bevel"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="bevel" d="M9.29 5.5h29.42a3.79 3.79 0 0 1 3.79 3.79v29.42a3.79 3.79 0 0 1-3.79 3.79H9.29a3.79 3.79 0 0 1-3.79-3.79V9.29A3.79 3.79 0 0 1 9.29 5.5Z"/><path fill="none" stroke="currentColor" d="M42.5 20.26h-7.68m-21.74.1H5.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="bevel" d="M30.63 14.68v-4a1.79 1.79 0 0 1 1.8-1.8h4.88a1.79 1.79 0 0 1 1.8 1.8v4.78a1.79 1.79 0 0 1-1.8 1.8h-4.18"/>`),
		g.Group(children),
	)
}