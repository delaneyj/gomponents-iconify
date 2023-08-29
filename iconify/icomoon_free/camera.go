package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.75 9.5a3.25 3.25 0 1 0 6.5 0a3.25 3.25 0 0 0-6.5 0zM15 4h-3.5c-.25-1-.5-2-1.5-2H6C5 2 4.75 3 4.5 4H1c-.55 0-1 .45-1 1v9c0 .55.45 1 1 1h14c.55 0 1-.45 1-1V5c0-.55-.45-1-1-1zm-7 9.938a4.438 4.438 0 1 1 0-8.876a4.438 4.438 0 0 1 0 8.876zM15 7h-2V6h2v1z"/>`),
		g.Group(children),
	)
}