package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scalable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 9.41v12.41L26.18 5.5h12.41a3.91 3.91 0 0 1 3.91 3.91ZM21.82 42.5H9.41a3.91 3.91 0 0 1-3.91-3.91V26.18L21.82 42.5Zm18.327-2.353h0a8.034 8.034 0 0 1-11.362 0L7.853 19.215a8.034 8.034 0 0 1 0-11.362h0a8.034 8.034 0 0 1 11.362 0l20.932 20.932a8.034 8.034 0 0 1 0 11.362Z"/>`),
		g.Group(children),
	)
}