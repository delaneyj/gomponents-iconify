package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 1.29v21.42L5.737 17.5H1v-11h4.737L15 1.29ZM4.999 8.5H3v7h1.999v-7Zm2 7.415L13 19.29V4.71L6.999 8.084v7.83ZM16 11h8v2h-8v-2Z"/>`),
		g.Group(children),
	)
}