package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 152v56a20 20 0 0 1-20 20H48a20 20 0 0 1-20-20v-56a12 12 0 0 1 24 0v52h152v-52a12 12 0 0 1 24 0ZM96.49 88.49L116 69v83a12 12 0 0 0 24 0V69l19.51 19.52a12 12 0 0 0 17-17l-40-40a12 12 0 0 0-17 0l-40 40a12 12 0 0 0 17 17Z"/>`),
		g.Group(children),
	)
}