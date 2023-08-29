package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M84.41 106c-15.63.1-27.67 13.8-25.69 29.3c16 124 16 117.4 0 241.4c-2.54 19.8 17.33 35 35.79 27.3L361.5 292.9v98.8c0 7.9 8.9 14.2 20 14.3h52c11.1-.1 20-6.4 20-14.3V120.2c-.1-7.8-9-14.1-20-14.2h-52c-11 .1-19.9 6.4-20 14.2v98.9L94.51 108c-3.2-1.3-6.63-2-10.1-2z"/>`),
		g.Group(children),
	)
}