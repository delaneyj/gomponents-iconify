package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Units(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28 34.3a2.63 2.63 0 0 1 2.63-2.63h0a2.62 2.62 0 0 1 2.62 2.63v4.2M28 31.67v6.83"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.25 34.3a2.63 2.63 0 0 1 2.63-2.63h0a2.63 2.63 0 0 1 2.63 2.63v4.2"/><circle cx="9.51" cy="10.39" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.51 13.61v6.95m15.4 16.62a2.64 2.64 0 0 1-2.28 1.32h0A2.63 2.63 0 0 1 20 35.87v-1.7a2.63 2.63 0 0 1 2.63-2.63h0a2.61 2.61 0 0 1 2.27 1.32m-7.04-12.3v-4.33a2.62 2.62 0 0 0-2.62-2.62h0a2.62 2.62 0 0 0-2.63 2.62v4.33m0-4.33v-2.62"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 35h-3a4.5 4.5 0 0 1-4.5-4.49h0a4.5 4.5 0 0 1 4.5-4.5h20a4.5 4.5 0 0 0 4.5-4.49h0a4.5 4.5 0 0 0-4.5-4.45H20.82"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.32 33.34l1.7 1.7l-1.7 1.7m7.2-17.97l-1.7-1.7l1.7-1.7"/>`),
		g.Group(children),
	)
}