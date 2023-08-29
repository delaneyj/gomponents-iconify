package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#F44336" d="M74.06 115.46c11.38 11.38 30.01 11.38 41.4 0c11.38-11.38 11.38-30.01 0-41.4L84.7 43.3L43.3 84.7l30.76 30.76z"/><path fill="#FDD835" d="M84.7 43.3L53.94 12.54c-11.38-11.38-30.01-11.38-41.4 0c-11.38 11.38-11.38 30.01 0 41.4L43.3 84.7l41.4-41.4z"/><path fill="#C62828" d="m86.819 45.423l-41.4 41.401l-2.122-2.12l41.4-41.403z"/><path fill="#FFF" d="M69.43 36.25C81.59 49.5 103.65 74.5 103.65 74.5c2.86 2.61 6.79 9.34 4.22 11.66c-2.74 2.48-9.12-2.75-11.98-5.36L49.22 36.89c-2.85-2.61-12.39-11.06-10.53-17.01c3.51-11.23 17.2 1.61 30.74 16.37z" opacity=".65"/>`),
		g.Group(children),
	)
}