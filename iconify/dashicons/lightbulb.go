package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 1c3.11 0 5.63 2.52 5.63 5.62c0 1.84-2.03 4.58-2.03 4.58c-.33.44-.6 1.25-.6 1.8v1c0 .55-.45 1-1 1H8c-.55 0-1-.45-1-1v-1c0-.55-.27-1.36-.6-1.8c0 0-2.02-2.74-2.02-4.58C4.38 3.52 6.89 1 10 1zM7 16.87V16h6v.87c0 .62-.13 1.13-.75 1.13H12c0 .62-.4 1-1.02 1h-2c-.61 0-.98-.38-.98-1h-.25c-.62 0-.75-.51-.75-1.13z"/>`),
		g.Group(children),
	)
}