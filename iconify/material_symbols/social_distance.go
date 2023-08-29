package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 21l-4-4l4-4l1.4 1.4L5.8 16h12.4l-1.6-1.6L18 13l4 4l-4 4l-1.4-1.4l1.6-1.6H5.8l1.6 1.6L6 21ZM2 11v-.575q0-.6.325-1.1t.9-.75q.65-.275 1.338-.425T6 8q.75 0 1.438.15t1.337.425q.575.25.9.75t.325 1.1V11H2Zm12 0v-.575q0-.6.325-1.1t.9-.75q.65-.275 1.338-.425T18 8q.75 0 1.438.15t1.337.425q.575.25.9.75t.325 1.1V11h-8ZM6 7q-.825 0-1.413-.588T4 5q0-.825.588-1.413T6 3q.825 0 1.413.588T8 5q0 .825-.588 1.413T6 7Zm12 0q-.825 0-1.413-.588T16 5q0-.825.588-1.413T18 3q.825 0 1.413.588T20 5q0 .825-.588 1.413T18 7Z"/>`),
		g.Group(children),
	)
}