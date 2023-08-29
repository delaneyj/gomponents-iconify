package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundMute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22.71L9.737 17.5H5v-11h4.737L19 1.29v21.42ZM8.999 8.5H7v7h1.999v-7Zm2 7.415L17 19.29V4.71l-6.001 3.375v7.83Z"/>`),
		g.Group(children),
	)
}