package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationDeskPeople(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 12.5h14.464m7.072 0H24M10 10V7.5s-1.5-1-4-1s-4 1-4 1V10m14 14v-5.5h-3.5v-.25l.072-.15A24.999 24.999 0 0 0 15 7.35v-.328a8.58 8.58 0 0 1 3-.523c1.288 0 2.311.266 3 .523v.328c0 3.72.83 7.391 2.428 10.749l.072.15v.25H20V24M5.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.746 3.5 6.15 4.5 6.15 4.5h-.3Zm12 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}