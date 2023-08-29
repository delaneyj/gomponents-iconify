package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playstation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M426 263q2 9-6.5 18T392 296l-2 1l-64-20l17-6q21-7 21-13q-2-10-37-4l-36 12l-61 21v22l96-32l64 20l-99 34l-61 21v-1v1l-69-22v-39v19q-40 14-84 6q-3 0-11-1.5t-12-2t-11-1.5t-11.5-2.5t-10-3t-8.5-3t-6.5-4t-5-5T0 288q-2-25 34-37l59 18l-15 6q-15 5-6 13q9 9 25 4l64-22v-44l-27-8l27-9V32l88 23q91 24 90 95q-1 90-82 67V100q0-6-7-9t-13.5-1t-6.5 9v148l6-2q58-20 104-17q80 6 86 35zM34 251h2l98-33l27 8v19l-68 24z"/>`),
		g.Group(children),
	)
}