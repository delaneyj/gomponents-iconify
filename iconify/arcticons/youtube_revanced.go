package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeRevanced(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.836 11.62l10.946 28.955h2.43l10.951-28.959"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.348 10.583l15.263-.008a1.215 1.215 0 0 1 1.073 1.786L24.952 26.88a1.081 1.081 0 0 1-1.907.002L15.278 12.37a1.214 1.214 0 0 1 1.07-1.786Z"/>`),
		g.Group(children),
	)
}