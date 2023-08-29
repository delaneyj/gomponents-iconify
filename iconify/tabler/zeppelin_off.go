package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeppelinOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.773 15.783c-.723.141-1.486.217-2.273.217c-2.13 0-4.584-.926-7.364-2.777L4 15v-3.33A46.07 46.07 0 0 1 2 10a46.07 46.07 0 0 1 2-1.67V5l2.135 1.778c.13-.087.261-.172.39-.256m2.564-1.42C10.69 4.367 12.16 4 13.5 4c4.694 0 8.5 2.686 8.5 6c0 1.919-1.276 3.627-3.261 4.725"/><path d="M10 15.5V20h6v-4M3 3l18 18"/></g>`),
		g.Group(children),
	)
}