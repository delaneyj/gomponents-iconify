package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolBoxOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469.3 85.3H42.7C19.1 85.3 0 104.5 0 128v21.3h512V128c0-23.5-19.1-42.7-42.7-42.7zm-256-42.6h85.3V64h42.7V42.7c0-23.5-19.1-42.7-42.7-42.7h-85.3c-23.5 0-42.7 19.1-42.7 42.7V64h42.7V42.7zM0 426.7c0 23.5 19.1 42.7 42.7 42.7h426.7c23.5 0 42.7-19.1 42.7-42.7v-21.3H0v21.3zm0-64h512V192H0v170.7z"/>`),
		g.Group(children),
	)
}