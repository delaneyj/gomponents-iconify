package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarthHero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 23.936c.038 10.91-9.378 19.649-20.543 18.455C13.471 41.483 6.59 34.65 5.624 26.17C4.353 15.014 13.024 5.538 23.936 5.5H42.5v18.436Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.602 23.97c.018 5.316-4.788 9.52-10.3 8.47c-3.355-.64-6.067-3.334-6.73-6.686c-1.089-5.503 3.084-10.338 8.399-10.356h8.63v8.573Z"/>`),
		g.Group(children),
	)
}