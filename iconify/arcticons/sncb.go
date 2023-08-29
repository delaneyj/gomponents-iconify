package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sncb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="19.5" ry="13.473"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.888 16.546v15.176m-3.646-13.734a14.602 14.602 0 0 1 8.478-2.459c3.476 0 5.596 1.993 5.596 4.07s-1.95 3.9-5.002 3.9h-5.426"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.242 29.857a13.628 13.628 0 0 0 8.083 2.798c3.702 0 7.432-1.4 7.432-4.748s-3.222-4.409-6.443-4.409"/>`),
		g.Group(children),
	)
}