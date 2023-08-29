package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cliqz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.387 18.146l4.19-1.402L12 12.027l4.716 11.578l1.403-4.19l3.917 3.917l1.268-1.268zm-7.387 1c.035 0 .07-.004.105-.004l1.908 4.686c-.654.11-1.326.172-2.013.172c-6.617 0-12-5.383-12-12S5.383 0 12 0s12 5.383 12 12c0 .695-.063 1.376-.177 2.04l-4.683-1.908c0-.044.006-.087.006-.133A7.153 7.153 0 0 0 12 4.854a7.155 7.154 0 0 0-7.147 7.145A7.155 7.154 0 0 0 12 19.146z"/>`),
		g.Group(children),
	)
}