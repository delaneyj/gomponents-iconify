package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DishwasherTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 3A3.25 3.25 0 0 0 3 6.25V10h1.5v-.5h15v8.25a1.75 1.75 0 0 1-1.75 1.75H7v.018c.766.11 1.373.716 1.482 1.482h9.268A3.25 3.25 0 0 0 21 17.75V6.25A3.25 3.25 0 0 0 17.75 3H6.25ZM19.5 8h-15V6.25c0-.966.784-1.75 1.75-1.75h11.5c.966 0 1.75.784 1.75 1.75V8ZM9 6.25a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3.75-.75a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5ZM2 11.75a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 .75.75v2A3.251 3.251 0 0 1 6 16.913V20.5h.75a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1 0-1.5h.75v-3.587A3.251 3.251 0 0 1 2 13.75v-2Z"/>`),
		g.Group(children),
	)
}