package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceNintendoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.713 4.718A4 4 0 0 0 3 8v8a4 4 0 0 0 4 4h3V10m0-4V4H8m6 6V4h3a4 4 0 0 1 4 4v8c0 .308-.035.608-.1.896m-1.62 2.39A3.982 3.982 0 0 1 17 20h-3v-6"/><path d="M5.5 8.5a1 1 0 1 0 2 0a1 1 0 1 0-2 0M3 3l18 18"/></g>`),
		g.Group(children),
	)
}