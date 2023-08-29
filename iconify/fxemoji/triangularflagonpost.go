package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangularflagonpost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#D5DEE4" d="M125.247 512h-11.653c-10.589 0-19.174-8.584-19.174-19.174V19.599C94.42 9.01 103.004.425 113.594.425h30.827v492.401c0 10.59-8.585 19.174-19.174 19.174z"/><path fill="#FF473E" d="M144.421.425v229.436l299.319-98.126c16.029-5.255 16.029-27.93 0-33.184L144.421.425z"/>`),
		g.Group(children),
	)
}