package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrophyThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 68h-28V56a12 12 0 0 0-12-12H64a12 12 0 0 0-12 12v12H24a12 12 0 0 0-12 12v16a36 36 0 0 0 36 36h6.66A76 76 0 0 0 124 187.89V220H96a4 4 0 0 0 0 8h64a4 4 0 0 0 0-8h-28v-32.12c32.93-1.74 60.41-24.91 69.11-55.88H208a36 36 0 0 0 36-36V80a12 12 0 0 0-12-12ZM48 124a28 28 0 0 1-28-28V80a4 4 0 0 1 4-4h28v36a77 77 0 0 0 1 12Zm148-12.9c0 37.71-30.27 68.62-67.48 68.9H128a68 68 0 0 1-68-68V56a4 4 0 0 1 4-4h128a4 4 0 0 1 4 4ZM236 96a28 28 0 0 1-28 28h-5.1a77.35 77.35 0 0 0 1.1-12.9V76h28a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}