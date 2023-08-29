package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phonepe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.805 8.5l9.092 10.331m-15.66 0h21.526M29.287 39.5c-.443-9.006-.39-20.669-.39-20.669"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.97 18.831c0 3.128-.119 7.93-.119 9.02a4.233 4.233 0 0 0 4.964 4.297a30.201 30.201 0 0 0 5.191-1.117"/>`),
		g.Group(children),
	)
}