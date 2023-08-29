package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwatchesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 155.91a16 16 0 0 0-1-5.22l-19.06-52.21A16 16 0 0 0 191.49 89l-67.81 24.57l12.08-69A16 16 0 0 0 122.84 26l-54.67-9.75a15.94 15.94 0 0 0-18.47 13l-25 143.12A43.82 43.82 0 0 0 67.78 224H216a16 16 0 0 0 16-16ZM68 196a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm42.72-8.38l9.78-55.92l76.42-27.7L216 156.11L108.78 195a44.89 44.89 0 0 0 1.94-7.38ZM216 208h-96.26L216 173.11Z"/>`),
		g.Group(children),
	)
}