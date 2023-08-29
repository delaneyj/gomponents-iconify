package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bunpo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.987 14.144a37.35 37.35 0 0 0 8.419 5.659a16.96 16.96 0 0 0 9.599 1.165a21.594 21.594 0 0 0 8.655-4.02a25.162 25.162 0 0 0 6.074-5.968C41 9.048 41.3 7.35 40.751 6.394s-1.948-1.167-3.614-.528a15.845 15.845 0 0 0-5.476 4.18a37.825 37.825 0 0 0-5.436 7.985a70.942 70.942 0 0 0-4.397 10.428a26.033 26.033 0 0 0-1.421 7.224a8.221 8.221 0 0 0 .85 4.491c.678 1.155 2.226 2.326 4.922 2.326c7.296 0 10.697-7.255 11.473-11.554c.506-2.804-.177-6.57-2.482-8.24"/>`),
		g.Group(children),
	)
}