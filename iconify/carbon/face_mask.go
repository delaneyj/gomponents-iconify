package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11.5 11a2.5 2.5 0 1 0 2.5 2.5a2.5 2.5 0 0 0-2.5-2.5zm9 0a2.5 2.5 0 1 0 2.5 2.5a2.5 2.5 0 0 0-2.5-2.5z"/><path fill="currentColor" d="M30 16A14 14 0 1 0 8.9 28.042l.015.156A1.993 1.993 0 0 0 10.905 30h10.19a1.993 1.993 0 0 0 1.99-1.8l.015-.158A13.967 13.967 0 0 0 30 16ZM16 4a11.99 11.99 0 0 1 11.97 11.397L22.764 18H9.236L4.03 15.397A11.99 11.99 0 0 1 16 4ZM4.12 17.678l3.94 1.97l.581 5.81a11.97 11.97 0 0 1-4.521-7.78ZM21.095 28h-10.19l-.8-8h11.79Zm2.264-2.542l.581-5.81l3.94-1.97a11.971 11.971 0 0 1-4.52 7.78Z"/>`),
		g.Group(children),
	)
}