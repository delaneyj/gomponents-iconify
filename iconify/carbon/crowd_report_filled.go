package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrowdReportFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="m16.832 19.555l-1.664-1.11L17.465 15H20v-5h-8v5h2v2h-2a2.002 2.002 0 0 1-2-2v-5a2.002 2.002 0 0 1 2-2h8a2.002 2.002 0 0 1 2 2v5a2.002 2.002 0 0 1-2 2h-1.465Z"/><path fill="currentColor" d="M16 2A11.013 11.013 0 0 0 5 13a10.889 10.889 0 0 0 2.216 6.6s.3.395.349.452L16 30l8.439-9.953c.044-.053.345-.447.345-.447l.001-.003A10.885 10.885 0 0 0 27 13A11.013 11.013 0 0 0 16 2Zm6 13a2.002 2.002 0 0 1-2 2h-1.465l-1.703 2.555l-1.664-1.11L17.465 15H20v-5h-8v5h2v2h-2a2.002 2.002 0 0 1-2-2v-5a2.002 2.002 0 0 1 2-2h8a2.002 2.002 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}