package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M79 0h164q16-2 29.5 9T287 37v111l36-59q30 47 84 135.5T474 335H233q-29 30-52 39q-42 18-88 4t-71-53q-27-39-21-89t41-82V48q-2-18 8.5-33T79 0zm12 53v72q40-13 82 1t65 49V53H91zm40 114q-30 1-51 18.5T49 233q-8 34 13 65.5t56 36.5q34 6 64-17t32-58q5-36-21-65t-62-28zm192 16l-62 102h125z"/>`),
		g.Group(children),
	)
}