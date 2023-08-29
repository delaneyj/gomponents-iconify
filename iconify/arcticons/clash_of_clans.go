package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClashOfClans(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.013 43.5s9.317-3.6 12.37-8.531s2.872-14.57 2.872-14.57a3.008 3.008 0 0 0-2.345-2.346c-1.848-.32-4.635-.744-4.635-.744a2.238 2.238 0 0 0-2.121-1.74a43.767 43.767 0 0 0-10.145-.357a2.324 2.324 0 0 0-1.752 1.53l-1.934.16a15.53 15.53 0 0 0-4.334.615a2.302 2.302 0 0 0-1.012 1.866c-.168 2.044-.345 7.868-.345 7.868a11.603 11.603 0 0 0 2.106 6.621c2.305 3.39 4.507 6.16 11.275 9.628Zm0 0l1.246-28.341m-2.319 7.862L10.424 12.153m4.756 3.109v-2.639L9.911 8.086l.511 4.068"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.18 15.262l-3.34.895l-6.097-3.316l4.68-.687m20.5 13.684L22.435 8.301m2.638 5.033l1.286-2.305L23.971 4.5l-1.536 3.801"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.073 13.334l-3.352-.846l-3.708-5.867l4.422 1.68"/>`),
		g.Group(children),
	)
}