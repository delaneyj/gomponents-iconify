package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M214 152a6 6 0 0 1-6 6h-74v59.51l21.76-21.75a6 6 0 0 1 8.48 8.48l-32 32a6 6 0 0 1-8.48 0l-32-32a6 6 0 0 1 8.48-8.48L122 217.51V158H48a6 6 0 0 1 0-12h160a6 6 0 0 1 6 6ZM48 110h160a6 6 0 0 0 0-12h-74V38.49l21.76 21.75a6 6 0 0 0 8.48-8.48l-32-32a6 6 0 0 0-8.48 0l-32 32a6 6 0 0 0 8.48 8.48L122 38.49V98H48a6 6 0 0 0 0 12Z"/>`),
		g.Group(children),
	)
}