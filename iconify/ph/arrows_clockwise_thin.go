package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsClockwiseThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M194.83 189.18a4 4 0 0 1 0 5.65c-1 1-25.65 25.17-66.83 25.17c-23.93 0-47.35-10.05-67.73-29.08a146.39 146.39 0 0 1-16.27-18V208a4 4 0 0 1-8 0v-48a4 4 0 0 1 4-4h48a4 4 0 0 1 0 8H47.41c10 14.06 38.39 48 80.59 48c37.75 0 60.95-22.6 61.18-22.83a4 4 0 0 1 5.65.01ZM216 44a4 4 0 0 0-4 4v35.07a146.39 146.39 0 0 0-16.27-18C175.35 46.05 151.93 36 128 36c-41.18 0-65.8 24.14-66.83 25.17a4 4 0 0 0 5.65 5.66C67.05 66.6 90.25 44 128 44c42.2 0 70.63 33.94 80.59 48H168a4 4 0 0 0 0 8h48a4 4 0 0 0 4-4V48a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}