package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsCounterClockwiseThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 100H40a4 4 0 0 1-4-4V48a4 4 0 0 1 8 0v35.07a146.39 146.39 0 0 1 16.27-18C80.65 46.05 104.07 36 128 36c41.18 0 65.8 24.14 66.83 25.17a4 4 0 1 1-5.66 5.66C188.75 66.41 165.58 44 128 44c-42.2 0-70.62 33.94-80.59 48H88a4 4 0 0 1 0 8Zm128 56h-48a4 4 0 0 0 0 8h40.59c-10 14.06-38.39 48-80.59 48c-37.58 0-60.75-22.41-61.17-22.83a4 4 0 0 0-5.66 5.66c1 1 25.65 25.17 66.83 25.17c23.93 0 47.35-10.05 67.73-29.08a146.39 146.39 0 0 0 16.27-18V208a4 4 0 0 0 8 0v-48a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}