package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etsy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Zm-31-14.55h5.96M9.5 16.02h5.96M9.5 21.99h3.89M9.5 16.02v11.93"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 25v4a3 3 0 0 1-3 3h0a3 3 0 0 1-2.11-.88"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 20.05V25a3 3 0 0 1-3 3h0a3 3 0 0 1-3-3v-4.95m-8.35 7.23a3.33 3.33 0 0 0 2.45.72h.67a2 2 0 0 0 2-2h0a2 2 0 0 0-2-2h-1.34a2 2 0 0 1-2-2h0a2 2 0 0 1 2-2h.67a3.37 3.37 0 0 1 2.46.67m-10.19-3.08v8.87A1.49 1.49 0 0 0 20.36 28h.44m-3.5-7.95h3.13"/>`),
		g.Group(children),
	)
}