package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHFiveLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M246 180a34 34 0 0 1-34 34a33.6 33.6 0 0 1-24.29-9.8a6 6 0 0 1 8.58-8.4A21.65 21.65 0 0 0 212 202a22 22 0 0 0 0-44a21.65 21.65 0 0 0-15.71 6.2a6 6 0 0 1-10.21-5.2l8-48a6 6 0 0 1 5.92-5h40a6 6 0 0 1 0 12h-34.92l-5 30a36 36 0 0 1 11.92-2a34 34 0 0 1 34 34ZM144 50a6 6 0 0 0-6 6v54H46V56a6 6 0 0 0-12 0v120a6 6 0 0 0 12 0v-54h92v54a6 6 0 0 0 12 0V56a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}