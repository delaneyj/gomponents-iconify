package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infinilist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.308 23.533l-1.934 2.284a2.043 2.043 0 0 1-2.667.352a2.07 2.07 0 0 1-.763-2.58a2.004 2.004 0 0 1 2.427-1.157l16.124 5.702a9.695 9.695 0 1 0-4.117-15.344l-6.483 7.68"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.872 26.922l-4.457 7.188a9.692 9.692 0 0 1-17.029-9.18a9.822 9.822 0 0 1 12.284-5.004l14.712 5.098a2.876 2.876 0 0 0 3.57-1.73a2.918 2.918 0 0 0-1.455-3.69a2.958 2.958 0 0 0-3.79 1.17l-1.856 2.98m8.694-2.075l4.597-.862m-18.367 2.415l-.526 1.186m5.354.471l-.551 1.19m5.069.174l-.112 1.366m3.671-2.73l1.493 1.464m-3.009-8.176l2.532-6.688m-7.078 7.47l-4.075-4.195m-2.919 14.643l-2.935-3.34m-.049 7.886l-4.392-7.057M15.85 37.21l-1.459-9.413M9.45 35.873l3.141-8.322m-6.635 3.202l5.212-4.258m-4.33-1.672l3.775-.057m.247-3.652l.61.923m3.58-1.572l-.353.97m4.673.225l-.561 1.18m13.795-4.766l-.858-6.897m4.886 8.544l4.664-4.344m-14.217 6.462l-3.46-3.596"/>`),
		g.Group(children),
	)
}