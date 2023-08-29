package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerraNil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.229 8.822A21.432 21.432 0 0 1 45.5 24c0 11.876-9.624 21.5-21.5 21.5S2.5 35.876 2.5 24S12.124 2.5 24 2.5a21.42 21.42 0 0 1 14.223 5.375"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.84 19.59a2.11 2.11 0 1 1 0 4.219a2.11 2.11 0 0 1 0-4.219Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.567 20.491S41.276 7.29 40.648 6.632s-10.792 7.613-11.952 9.035c-1.897 2.326-.127 4.824-.127 4.824Zm-3.257-.242s-15.482-9.803-16-9.055s9.622 9.046 11.248 9.897c2.659 1.392 4.752-.842 4.752-.842Zm1.629 3.556s4.38 17.793 5.272 17.613s-.677-13.19-1.261-14.93c-.954-2.845-4.011-2.682-4.011-2.682Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.891 21.504c-.83 1.555-.4 2.702 1.058 3.131"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.375 45.49l.573-21.686c.54-.39 1.519-.503 1.519-.503m.655.503l.817 21.497"/>`),
		g.Group(children),
	)
}