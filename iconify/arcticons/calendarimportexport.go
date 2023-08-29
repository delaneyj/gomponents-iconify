package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendarimportexport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 28.6a8.46 8.46 0 1 0 8.46 8.46h0A8.45 8.45 0 0 0 37 28.6Zm-2.52 4.23v8.45m-2.75-5.56l2.75-2.89m2.74 2.89l-2.74-2.89m5.08 8.45v-8.45m2.75 5.57l-2.75 2.88m-2.74-2.88l2.74 2.88M15.8 8.58a2.51 2.51 0 1 1-2.51 2.51h0a2.51 2.51 0 0 1 2.51-2.51Zm16.4 0a2.51 2.51 0 0 1 2.51 2.51h0a2.51 2.51 0 0 1-2.51 2.51h0a2.51 2.51 0 1 1 0-5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 30.63V7.45a2 2 0 0 0-2-1.95H7.45a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h23.02"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".96" d="m27.75 22.22l4.5-2.45m0 0v10.29m-13.88 7.71a4.5 4.5 0 0 0 4.5-4.5h0a4.49 4.49 0 0 0-4.5-4.5h0a4.5 4.5 0 0 0 4.5-4.5h0a4.5 4.5 0 0 0-4.5-4.5m-7.12 16.48c1.24 1 2.58 1.52 5.6 1.52h1.52m-7.12-16.5c1.24-1 2.59-1.51 5.6-1.5h1.52"/>`),
		g.Group(children),
	)
}