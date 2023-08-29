package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignpostROne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.563 8h13.9a2.489 2.489 0 0 1 1.744.709l2.25 2.192a1.5 1.5 0 0 1-.007 2.2l-2.243 2.187a2.483 2.483 0 0 1-1.743.712H3.563a1.5 1.5 0 0 1-1.5-1.5v-5a1.5 1.5 0 0 1 1.5-1.5Zm13.9 7a1.491 1.491 0 0 0 1.046-.425l2.255-2.2a.5.5 0 0 0 .173-.375a.494.494 0 0 0-.162-.369l-.01-.01l-2.254-2.2A1.492 1.492 0 0 0 17.464 9H3.563a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5Z"/>`),
		g.Group(children),
	)
}