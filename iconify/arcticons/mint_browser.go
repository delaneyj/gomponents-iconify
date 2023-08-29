package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MintBrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.466 8.094c7.117 14.179-2.83 22.36-11.14 24.2m-16.375 8.644c-8.368-14.918 1.82-23.53 10.36-25.47m-3.28-12.12c23.828 2.069 16.49 28.311 9.222 29.019m3.443 12.026c-15.246-1.325-17.73-12.55-15.636-20.646M2.786 20.605c8.733-14.715 21.763-9.595 27.466-3.01m15.05 9.24c-8.11 13.91-20.006 10.478-26.254 4.62M32.95 24A8.95 8.95 0 0 1 24 32.95h0A8.95 8.95 0 0 1 15.05 24h0A8.95 8.95 0 0 1 24 15.05h0A8.95 8.95 0 0 1 32.95 24h0Zm12.55 0c0 11.874-9.626 21.5-21.5 21.5S2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24Z"/>`),
		g.Group(children),
	)
}