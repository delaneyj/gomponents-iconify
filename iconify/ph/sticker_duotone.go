package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickerDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 136c-8 24-56 72-80 80v-32a48 48 0 0 1 48-48Z" opacity=".2"/><path d="M168 32H88a56.06 56.06 0 0 0-56 56v80a56.06 56.06 0 0 0 56 56h48a8.07 8.07 0 0 0 2.53-.41c26.23-8.75 76.31-58.83 85.06-85.06A8.07 8.07 0 0 0 224 136V88a56.06 56.06 0 0 0-56-56ZM48 168V88a40 40 0 0 1 40-40h80a40 40 0 0 1 40 40v40h-24a56.06 56.06 0 0 0-56 56v24H88a40 40 0 0 1-40-40Zm96 35.14V184a40 40 0 0 1 40-40h19.14C191 163.5 163.5 191 144 203.14Z"/></g>`),
		g.Group(children),
	)
}