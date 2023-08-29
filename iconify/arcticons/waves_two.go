package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WavesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 26c-4.875 0-4.875-4-9.749-4s-4.873 4-9.746 4c-4.876 0-4.876-4-9.752-4S9.375 26 4.5 26m39 10c-4.875 0-4.875-4-9.749-4s-4.873 4-9.746 4c-4.876 0-4.876-4-9.752-4S9.375 36 4.5 36m39-20c-4.875 0-4.875-4-9.749-4s-4.873 4-9.746 4c-4.876 0-4.876-4-9.752-4S9.375 16 4.5 16"/>`),
		g.Group(children),
	)
}