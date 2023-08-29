package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 149.3V448c0 23.5 19.1 42.7 42.7 42.7H64v-384H42.7C19.1 106.7 0 125.8 0 149.3zm469.3-42.6H448v384h21.3c23.5 0 42.7-19.1 42.7-42.7V149.3c0-23.5-19.1-42.6-42.7-42.6zm-362.6 384h298.7v-384H106.7v384zM213.3 64h85.3v21.3h42.7V64c0-23.5-19.1-42.7-42.7-42.7h-85.3c-23.5 0-42.7 19.1-42.7 42.7v21.3h42.7V64z"/>`),
		g.Group(children),
	)
}