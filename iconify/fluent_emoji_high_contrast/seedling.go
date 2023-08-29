package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seedling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.61 2a7.156 7.156 0 0 1 7.14 6.73h.01v.47l.04-.05c.115-.145.236-.285.362-.42l.024-.026a7.257 7.257 0 0 1 .574-.546a7.113 7.113 0 0 1 4.63-1.708v.01H30a7.15 7.15 0 0 1-7.15 7.15h-6.09v7.47c4.401.59 7.79 4.355 7.79 8.91H6.56c0-4.202 2.884-7.731 6.78-8.716V9.15H9.15A7.15 7.15 0 0 1 2 2h7.61Zm6.15 19.002V12.61h7.09a6.15 6.15 0 0 0 6.07-5.15h-6.883a6.142 6.142 0 0 0-4.49 2.357l-1.787 2.278V8.93l-.008-.141A6.155 6.155 0 0 0 9.61 3H3.08a6.15 6.15 0 0 0 6.07 5.15h5.19v12.93a9.064 9.064 0 0 1 1.42-.078Z"/>`),
		g.Group(children),
	)
}