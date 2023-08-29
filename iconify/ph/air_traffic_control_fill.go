package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirTrafficControlFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M229.11 62.82A16 16 0 0 0 216 56h-80V24h16a8 8 0 0 0 0-16h-48a8 8 0 0 0 0 16h16v32H40a16 16 0 0 0-15 21.47l26.19 72a16 16 0 0 0 15 10.53h123.63a16 16 0 0 0 15-10.53l26.19-72a16 16 0 0 0-1.9-14.65Zm-126.59 81.05a7.87 7.87 0 0 1-1.44.13a8 8 0 0 1-7.86-6.57L83 81.43a8 8 0 0 1 15.75-2.86l10.18 56a8 8 0 0 1-6.41 9.3ZM173 81.43l-10.19 56a8 8 0 0 1-7.86 6.57a7.87 7.87 0 0 1-1.44-.13a8 8 0 0 1-6.44-9.3l10.18-56A8 8 0 0 1 173 81.43ZM160 180v44a8 8 0 0 1-8 8h-48a8 8 0 0 1-8-8v-44a4 4 0 0 1 4-4h56a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}