package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundNightlightRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 22h.21c.84-.02 1.12-1.11.41-1.56a9.99 9.99 0 0 1-4.63-8.43c0-3.55 1.85-6.66 4.63-8.44c.7-.45.44-1.54-.39-1.56h-.13c-4.9-.05-9.21 3.53-9.98 8.37C4.64 16.61 9.45 22 15.5 22z"/>`),
		g.Group(children),
	)
}