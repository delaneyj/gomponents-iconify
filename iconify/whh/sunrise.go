package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 759q-97 62-96 137q1 36 4 55q92 33 92 73H0q0-40 93-74q3-27 3-54q-2-77-96-137q255-11 137-237q227 117 238-138q60 94 137 96q75 2 137-96q11 255 238 138q-118 226 137 237zM512 576q-87 0-161 43T234.5 735.5T192 896q0 12 1 28q30-6 64-11q-1-15-1-17q0-106 75-181t181-75t181 75t75 181q0 3-1 17q34 5 64 11q1-17 1-28q0-87-42.5-160.5T673 619t-161-43zm64-384v96q0 13-9.5 22.5T544 320h-64q-13 0-22.5-9.5T448 288v-96H335q-25-22-8-37L489 9q9-9 23-9t24 9l162 146q16 15-9 37H576z"/>`),
		g.Group(children),
	)
}