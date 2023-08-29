package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToiletEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M3.33 2.19a1.1 1.1 0 1 1 1.1-1.1a1.1 1.1 0 0 1-1.1 1.1zm6.94-1.1a1.1 1.1 0 1 0 0 .01v-.01zM6.51 4.93L4.7 3.12A.37.37 0 0 0 4.43 3H2.22a.37.37 0 0 0-.25.1H2L.14 4.93a.38.38 0 1 0 .53.53l1.58-1.58L.77 8h1.46v2.51a.38.38 0 0 0 .75.11H3V8h.69v2.63a.38.38 0 0 0 .75-.11V8h1.44L4.41 3.88L6 5.46a.37.37 0 0 0 .28.12c.21 0 .38-.17.38-.38a.37.37 0 0 0-.15-.27zM8.62 7v3.63a.37.37 0 1 0 .73 0V7H11V3.37a.37.37 0 0 0-.37-.37H7.71a.37.37 0 0 0-.37.37V7h1.28z" fill="currentColor"/>`),
		g.Group(children),
	)
}