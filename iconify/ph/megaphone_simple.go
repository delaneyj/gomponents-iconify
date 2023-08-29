package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m220.54 86.66l-176.06-54A16 16 0 0 0 24 48v144a16 16 0 0 0 16 16a16 16 0 0 0 4.52-.65L128 181.73V192a16 16 0 0 0 16 16h32a16 16 0 0 0 16-16v-29.9l28.54-8.75A16.09 16.09 0 0 0 232 138v-36a16.09 16.09 0 0 0-11.46-15.34ZM128 165l-88 27V48l88 27Zm48 27h-32v-15.18l32-9.82Zm40-54h-.11L144 160.08V79.92l71.89 22h.11v36Z"/>`),
		g.Group(children),
	)
}