package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.94 14.74a17.38 17.38 0 0 1 14.7.17A17 17 0 0 1 39.17 22l-3.25 1.57A15.94 15.94 0 0 0 15 17a17 17 0 0 0-5.74 4.82a17.28 17.28 0 0 1 7.66-7.11Zm5.11 5.8a11.28 11.28 0 0 0-8.35 3.66A12.46 12.46 0 0 1 24 18.36a12.1 12.1 0 0 1 10.31 6l-3.15 1.54a11 11 0 0 0-9.11-5.36Zm7.52 6.14l-3.16 1.58a5.8 5.8 0 0 0-4.33-2.8a5.74 5.74 0 0 0-3.63 1.08a7.08 7.08 0 0 1 5.06-3.25a6.59 6.59 0 0 1 6.06 3.39Zm-5.76 8.08L4.74 25.22a.43.43 0 0 1-.24-.39v-2.16a.22.22 0 0 1 .31-.2l19 9.5a.37.37 0 0 0 .38 0l19-9.49a.22.22 0 0 1 .31.2v2.16a.43.43 0 0 1-.24.39l-19.07 9.54a.39.39 0 0 1-.38 0Zm0 0"/>`),
		g.Group(children),
	)
}