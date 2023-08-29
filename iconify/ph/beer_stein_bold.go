package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeerSteinBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 84h-12V72a44.05 44.05 0 0 0-44-44h-9.73c-12.5-10.22-29.09-16-46.27-16c-37.5 0-68 26.92-68 60v136a20 20 0 0 0 20 20h128a20 20 0 0 0 20-20v-4h12a28 28 0 0 0 28-28v-64a28 28 0 0 0-28-28ZM104 36c12.85 0 25 4.62 33.44 12.67a12 12 0 0 0 8.3 3.33H160a20 20 0 0 1 19.6 16H60.28C62.72 50 81.39 36 104 36Zm76 168H60V92h120Zm40-28a4 4 0 0 1-4 4h-12v-72h12a4 4 0 0 1 4 4Zm-112-56v56a12 12 0 0 1-24 0v-56a12 12 0 0 1 24 0Zm48 0v56a12 12 0 0 1-24 0v-56a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}