package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Computer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="m18.5 35.5l-8 2c-.48.04-1 .52-1 1v1c0 .5.47 1 1 1h20c.43 0 1-.41 1-1v-1c-.02-.52-.55-.98-1-1l-8-2h-4zm19-1c2.59 0 3-.529 3-3v-26c0-2.391-.55-3-3-3h-34c-2.43 0-3 .54-3 3v26c0 2.51.529 3 3 3h34zm-2-27v22h-30v-22h30z"/>`),
		g.Group(children),
	)
}