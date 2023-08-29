package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundMuteOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 1.29v21.42L5.737 17.5H1v-11h4.737L15 1.29ZM4.999 8.5H3v7h1.999v-7Zm2 7.415L13 19.29V4.71L6.999 8.084v7.83Zm10.88-7.45L20 10.584l2.121-2.12l1.415 1.413L21.413 12l2.121 2.121l-1.414 1.414L20 13.414l-2.121 2.121l-1.415-1.414L18.587 12l-2.121-2.122l1.414-1.414Z"/>`),
		g.Group(children),
	)
}