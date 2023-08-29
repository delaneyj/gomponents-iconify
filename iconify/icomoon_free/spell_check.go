package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpellCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 4h2v3h1V1c0-.55-.45-1-1-1H2c-.55 0-1 .45-1 1v6h1V4zm0-3h2v2H2V1zm13 0V0h-3c-.55 0-1 .45-1 1v5c0 .55.45 1 1 1h3V6h-3V1h3zm-5 1.5V1c0-.55-.45-1-1-1H6v7h3c.55 0 1-.45 1-1V4.5c0-.55-.137-1-.688-1c.55 0 .688-.45.688-1zM9 6H7V4h2v2zm0-3H7V1h2v2zm4 6l-6.5 7L3 11.5l1.281-1.094L6.5 12.719L12 8z"/>`),
		g.Group(children),
	)
}