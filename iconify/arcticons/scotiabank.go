package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scotiabank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.5 12l6.749-7.5h-19.5c-9 0-15 11.25-9.75 18C12 18 18 12 21 12h13.5Zm-21 24l-6.749 7.5H26.25c9 0 14.999-11.25 9.75-18C36 30 30 36 27 36H13.5Z"/><circle cx="24" cy="24" r="9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}