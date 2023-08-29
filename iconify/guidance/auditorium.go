package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Auditorium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M11 19.5h7m-.5-5.5V6.5h-1.172a3 3 0 0 0-2.906 2.255L11.5 16.25v.25h9V18c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5m-21-4h7M8.5 14V6.5H7.328a3 3 0 0 0-2.906 2.255L2.5 16.25v.25H9m8.35-12s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Zm-9 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}