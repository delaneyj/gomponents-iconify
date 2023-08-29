package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OtpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3.5 6a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 1 0v-2a.5.5 0 0 0-.5-.5ZM11 7h.5a.5.5 0 0 0 0-1H11v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-12Zm2 5a1.5 1.5 0 1 1 3 0v2a1.5 1.5 0 1 1-3 0v-2ZM7 6H6V5h3v1H8v4H7V6Zm3-1h1.5a1.5 1.5 0 0 1 0 3H11v2h-1V5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}