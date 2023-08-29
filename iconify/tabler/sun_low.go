package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12a4 4 0 1 0 8 0a4 4 0 1 0-8 0m-4 0h.01M12 4v.01M20 12h.01M12 20v.01M6.31 6.31L6.3 6.3m11.41.01l-.01-.01m0 11.4l.01.01M6.3 17.7l.01.01"/>`),
		g.Group(children),
	)
}