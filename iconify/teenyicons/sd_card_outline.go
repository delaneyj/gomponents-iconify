package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6.5 3v3m2-3v3m2-3v3m-8 8.5h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-7l-4 4v9a1 1 0 0 0 1 1Z"/>`),
		g.Group(children),
	)
}