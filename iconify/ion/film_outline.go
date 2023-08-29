package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="416" height="320" x="48" y="96" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="80" height="80" x="384" y="336" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="80" height="80" x="384" y="256" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="80" height="80" x="384" y="176" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="80" height="80" x="384" y="96" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="80" height="80" x="48" y="336" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="80" height="80" x="48" y="256" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="80" height="80" x="48" y="176" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="80" height="80" x="48" y="96" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="256" height="160" x="128" y="96" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/><rect width="256" height="160" x="128" y="256" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="28" ry="28"/>`),
		g.Group(children),
	)
}