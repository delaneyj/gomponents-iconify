package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnifeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M234.64 29.37a32 32 0 0 0-45.26 0l-.1.1L15.41 207.79a12 12 0 0 0 5.89 20.1A159.13 159.13 0 0 0 57.19 232c34.21 0 68.42-11.11 100.55-32.9c32.29-21.9 51.09-46.74 51.87-47.78a12 12 0 0 0-1.11-15.7l-16-16l42.29-45.08a32.09 32.09 0 0 0-.15-45.17Zm-91 150.25C112.74 200.43 81 209.85 49 207.77l97.6-100.13l37.13 37.16a222.39 222.39 0 0 1-40.06 34.82Zm74-121.94l-.26.27l-41.89 44.66l-12.12-12.17l43-44.12a8 8 0 0 1 11.28 11.36Z"/>`),
		g.Group(children),
	)
}