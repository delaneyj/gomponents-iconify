package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Offlinecalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 28.6a8.46 8.46 0 1 0 8.46 8.46h0A8.45 8.45 0 0 0 37 28.6Zm3.43 11.76l-6.23.09a3 3 0 0 1-3-3a3 3 0 0 1 3.1-3.12L33 33l8.76 8.65Zm-6.08-6h0a3.25 3.25 0 0 1 5.69.09a4.8 4.8 0 0 1 .43 1a2.94 2.94 0 0 1 2.53 2.8a1.89 1.89 0 0 1-1.8 2.11h-.74M15.8 8.58a2.51 2.51 0 1 1-2.51 2.51h0a2.51 2.51 0 0 1 2.51-2.51Zm16.4 0a2.51 2.51 0 0 1 2.51 2.51h0a2.51 2.51 0 0 1-2.51 2.51h0a2.51 2.51 0 1 1 0-5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 30.63V7.45a2 2 0 0 0-2-1.95H7.45a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h23.02"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".96" d="m27.75 22.22l4.5-2.45m0 0v10.29m-13.88 7.71a4.5 4.5 0 0 0 4.5-4.5h0a4.49 4.49 0 0 0-4.5-4.5h0a4.5 4.5 0 0 0 4.5-4.5h0a4.5 4.5 0 0 0-4.5-4.5m-7.12 16.48c1.24 1 2.58 1.52 5.6 1.52h1.52m-7.12-16.5c1.24-1 2.59-1.51 5.6-1.5h1.52"/>`),
		g.Group(children),
	)
}