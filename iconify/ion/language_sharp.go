package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M363 176L246 464h47.24l24.49-58h90.54l24.49 58H480Zm-26.69 186L363 279.85L389.69 362ZM272 320c-.25-.19-20.59-15.77-45.42-42.67c39.58-53.64 62-114.61 71.15-143.33H352V90H214V48h-44v42H32v44h219.25c-9.52 26.95-27.05 69.5-53.79 108.36c-32.68-43.44-47.14-75.88-47.33-76.22L143 152l-38 22l6.87 13.86c.89 1.56 17.19 37.9 54.71 86.57c.92 1.21 1.85 2.39 2.78 3.57c-49.72 56.86-89.15 79.09-89.66 79.47L64 368l23 36l19.3-11.47c2.2-1.67 41.33-24 92-80.78c24.52 26.28 43.22 40.83 44.3 41.67L255 362Z"/>`),
		g.Group(children),
	)
}