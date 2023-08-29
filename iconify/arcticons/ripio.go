package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ripio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.693 27.848l-13.225 6.093a.63.63 0 0 1-.895-.572V21.298a1.262 1.262 0 0 1 .734-1.146l13.225-6.093a.63.63 0 0 1 .895.572v12.071a1.262 1.262 0 0 1-.734 1.146Z"/>`),
		g.Group(children),
	)
}