package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragonZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="m34.021 42.494l3.74-3.74a6 6 0 0 0 0-8.485v0a6 6 0 0 0-8.485 0L27.045 32.5m-9.97-7l6.544-6.544a6 6 0 0 0 0-8.486v0a6 6 0 0 0-8.485 0l-7.071 7.071m9.012 7.959L8.77 33.806a6 6 0 0 0 0 8.485v0a6 6 0 0 0 8.485 0l9.766-9.766"/><path stroke-linejoin="round" d="M13 12V4m25 26l5-5"/></g>`),
		g.Group(children),
	)
}