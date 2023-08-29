package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Librus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 33.26v2.02c0 3.99-3.23 7.22-7.22 7.22H12.72c-3.99 0-7.22-3.23-7.22-7.22V12.72c0-3.99 3.23-7.22 7.22-7.22H24v20.49a7.27 7.27 0 0 0 7.27 7.27H42.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 12.72v11.43h-9.14V5.5h1.92c3.99 0 7.22 3.23 7.22 7.22Z"/>`),
		g.Group(children),
	)
}