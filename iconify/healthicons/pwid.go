package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pwid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m33 14.414l1.293 1.293l1.414-1.414l-4-4l-1.414 1.414L31.586 13L28.5 16.086l-2.793-2.793l-1.414 1.414L25.586 16l-5.879 5.879a3.001 3.001 0 0 0-.529 3.529l-1.31 1.31l-2.164-2.454l-7.377 7.765l4.124 2.886A6 6 0 0 0 15.89 36H31a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2h-.5v6h-3a3 3 0 0 1-3-3v-1.5h2V31a1 1 0 0 0 1 1h1v-4a2 2 0 0 0-2-2H34l-3 2h-9.464l-4 4h-2.122l5.179-5.178a3.001 3.001 0 0 0 3.528-.53L30 20.415l1.293 1.293l1.414-1.414l-2.793-2.793L33 14.414Zm-9.021 6.022L27 17.414L28.586 19l-3.022 3.021l-1.585-1.585ZM16.475 33.06l-.06-.061h.12l-.06.06Z" clip-rule="evenodd"/><path d="m4 11l10.375 11.759l-7.705 8.11L4 29V11Z"/></g>`),
		g.Group(children),
	)
}