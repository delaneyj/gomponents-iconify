package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opensource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M692 991L583 690q54-21 87.5-69.5T704 512q0-80-56-136t-136-56t-136 56t-56 136q0 60 33.5 108.5T441 690L332 991Q185 935 92.5 804.5T0 512q0-104 40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199q0 162-92.5 292.5T692 991z"/>`),
		g.Group(children),
	)
}