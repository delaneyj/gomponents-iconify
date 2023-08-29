package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Librerapro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.45v35.1a2 2 0 0 0 1.95 2h2.38V4.5h-2.38A2 2 0 0 0 8.4 6.45ZM37.65 4.5H34v15.6l-2.92-2l-2.93 2V4.5H12.73v39h24.92a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95Z"/>`),
		g.Group(children),
	)
}