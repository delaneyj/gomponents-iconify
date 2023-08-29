package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 19C44 12.9249 39.0751 8 33 8C29.2797 8 25.9907 9.8469 24 12.6738C22.0093 9.8469 18.7203 8 15 8C8.92487 8 4 12.9249 4 19C4 30 17 40 24 42.3262C25.1936 41.9295 26.5616 41.3098 28.0099 40.5"/><path d="M34.959 27L41.8375 35.5"/><path d="M41.8375 27L34.959 35.5"/></g>`),
		g.Group(children),
	)
}