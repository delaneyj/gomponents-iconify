package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnknownStatusLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="17.58" cy="26.23" r="1.4" fill="currentColor" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M24.7 13a5.18 5.18 0 0 0-2.16-3.56a7.26 7.26 0 0 0-5.71-1.09A11.34 11.34 0 0 0 12 10.44A1 1 0 1 0 13.26 12a9.32 9.32 0 0 1 3.94-1.72a5.29 5.29 0 0 1 4.16.74a3.21 3.21 0 0 1 1.35 2.19c.33 2.69-3.19 3.75-5.32 4.14l-.82.15v4.36a1 1 0 0 0 2 0v-2.69c6.04-1.38 6.31-4.76 6.13-6.17Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}