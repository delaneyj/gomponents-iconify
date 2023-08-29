package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserTransmit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm12.578 5.596l3.674 2.88v1.526H13.75v-2h4.655l-1.061-.832l1.234-1.574ZM8 16a4 4 0 0 0-4 4h7.8v2H2v-2a6 6 0 0 1 6-6h3.75v2H8Zm5.752 2.498h8.502v2H17.6l1.061.832l-1.234 1.574l-3.674-2.88v-1.526Z"/>`),
		g.Group(children),
	)
}