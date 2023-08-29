package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickerLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 34H88a54.06 54.06 0 0 0-54 54v80a54.06 54.06 0 0 0 54 54h48a5.86 5.86 0 0 0 1.9-.31c25.84-8.61 75.18-57.95 83.79-83.79a5.86 5.86 0 0 0 .31-1.9V88a54.06 54.06 0 0 0-54-54ZM46 168V88a42 42 0 0 1 42-42h80a42 42 0 0 1 42 42v42h-26a54.06 54.06 0 0 0-54 54v26H88a42 42 0 0 1-42-42Zm96 38.67V184a42 42 0 0 1 42-42h22.67c-11.83 21.1-43.57 52.84-64.67 64.67Z"/>`),
		g.Group(children),
	)
}