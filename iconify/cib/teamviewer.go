package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Teamviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30.13 32H1.875A1.88 1.88 0 0 1 0 30.125V1.875A1.88 1.88 0 0 1 1.875 0H30.13a1.884 1.884 0 0 1 1.875 1.875v28.25A1.884 1.884 0 0 1 30.13 32zM15.88 2.813C8.677 2.876 2.864 8.787 2.812 16.001C2.76 23.345 8.651 29.178 16 29.189c7.219 0 13.063-6 13.188-13.188c.125-7.432-5.875-13.255-13.307-13.188zm-3.015 8.697l-1.182 2.88h8.719l-1.188-2.88l8.859 4.484l-8.859 4.49l1.188-2.88h-8.719l1.177 2.88l-8.88-4.49z"/>`),
		g.Group(children),
	)
}