package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Codecheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.11 32.54s-3.89-3.09-5.84-4.86a79.18 79.18 0 0 1-5.67-6.44c1-1.36 5.79-5.73 5.79-5.73a26.89 26.89 0 0 1 2.71 2.25c1.07 1.06 3.67 4.24 3.67 4.24c1.94-1.41 5.23-6.49 5.23-6.49c-.92-1.06-5-4.72-5-4.72a51.55 51.55 0 0 0-13.48 0c-.52.35-3.29 3.31-3.29 3.31C4 17.39 4.7 28.21 4.7 28.21c.94 1.42 9 8.86 9 8.86s6.21.59 8 .59s6.86-.66 6.86-.66a43.06 43.06 0 0 0 7.56-7.09a61.4 61.4 0 0 0 7.38-11.1l-6.26-5.49c-3.6 7.5-5.85 9.91-8.45 12.58a56.86 56.86 0 0 1-15.06 11.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.7 28.21a49.8 49.8 0 0 1 4.9-7m5.79-5.7A72.88 72.88 0 0 1 22 10.79"/>`),
		g.Group(children),
	)
}