package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M471.05 168.36L263.24 65.69a16.37 16.37 0 0 0-14.48 0L41 168.36a16 16 0 0 0-9 14.31V432a16.09 16.09 0 0 0 16.19 16h415.62A16.09 16.09 0 0 0 480 432V182.67a16 16 0 0 0-8.95-14.31ZM256 97.89l173 85.44l-175.7 86.78l-173-85.44Z"/>`),
		g.Group(children),
	)
}