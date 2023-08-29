package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediumFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.001 3h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm13.3 12.94c-.1-.05-.15-.2-.15-.301V8.006c0-.1.05-.25.15-.351l.955-1.105V6.5h-3.415l-2.56 6.478L9.367 6.5H5.852v.05l.904 1.256c.201.2.251.502.251.753v5.523c.05.302 0 .653-.15.954L5.5 16.894v.05h3.616v-.05L7.76 15.087c-.15-.302-.201-.603-.15-.954V9.11c.05.1.1.1.15.301l3.414 7.633h.05L14.54 8.76c-.05.3-.05.652-.05.904v5.925c0 .15-.05.25-.15.351l-1.005.954v.05h4.921v-.05l-.954-.954Z"/>`),
		g.Group(children),
	)
}