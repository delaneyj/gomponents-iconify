package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Affinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.368 1.08h3.778l.318.55h1.082L24 18.004v.001l-2.036 3.47H13.69l.84 1.445h-.365l-.84-1.446H3.057l-.526-.923h-.652L0 17.298l.002-.001l2.41-4.176l2.23-1.288l3.69-6.39l-.742-1.285L9.368 1.08zm2.224 5.652L5.066 18.008h6.25l-.723-1.246l6.808.006l-5.809-10.036Z"/>`),
		g.Group(children),
	)
}