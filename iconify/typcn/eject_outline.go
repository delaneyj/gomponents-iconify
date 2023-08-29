package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 8.551V7a2 2 0 0 0-2-2H4v10a2 2 0 0 0 2 2h1.066c.893 2.887 3.646 5 6.809 5c3.859 0 7.062-3.016 7.062-6.875c.001-3.161-2.068-5.74-4.937-6.574zm-8 1.862v3.243a1 1 0 0 1-2 0V7h6.656a1 1 0 0 1 0 2H9.413l4.801 4.799a1 1 0 0 1-.707 1.707a.996.996 0 0 1-.707-.291L8 10.413zm6 9.618c-2.757 0-5-2.26-5-5.016c0-.705-.004-1.371.21-1.979l2.883 2.884c.39.39.901.584 1.414.584s1.022-.194 1.414-.584a2.002 2.002 0 0 0 0-2.83l-2.567-2.567c.517-.226 1.065-.398 1.646-.398c2.757 0 5 2.182 5 4.938c0 2.757-2.243 4.968-5 4.968z"/>`),
		g.Group(children),
	)
}