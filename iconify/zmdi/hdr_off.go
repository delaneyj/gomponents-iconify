package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M343 272h-8l-24-23V144h75q13 0 22.5 9.5T418 176v21q0 10-5.5 18.5T399 227l19 45h-32l-19-43h-24v43zm0-96v21h43v-21h-43zm-96 0h-8l-32-32h40q13 0 22.5 9.5T279 176v41l-32-32v-9zm-74-21l258 256l-24 23l-162-162h-72v-73l-32-32v105h-32v-53H66v53H34V144h32v43h43v-43h8L0 27L23 5z"/>`),
		g.Group(children),
	)
}