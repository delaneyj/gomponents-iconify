package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 1.29v21.42L5.737 17.5H1v-11h4.737L15 1.29ZM4.999 8.5H3v7h1.999v-7Zm2 7.415L13 19.29V4.71L6.999 8.084v7.83ZM19 8h2v3h3v2h-3v3h-2v-3h-3v-2h3V8Z"/>`),
		g.Group(children),
	)
}