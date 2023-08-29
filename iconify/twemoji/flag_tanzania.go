package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTanzania(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#1EB53A" d="M4 5a4 4 0 0 0-4 4v15.627L26.456 5H4z"/><path fill="#00A3DD" d="M32 31a4 4 0 0 0 4-4V11.33L9.479 31H32z"/><path fill="#141414" d="M32 5h-2.532L0 26.638V27a4 4 0 0 0 4 4h2.467L36 9.318V9a4 4 0 0 0-4-4z"/><path fill="#FBD035" d="M26.456 5L0 24.627v2.011L29.468 5zM9.479 31L36 11.33V9.318L6.467 31z"/>`),
		g.Group(children),
	)
}