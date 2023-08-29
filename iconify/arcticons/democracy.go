package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Democracy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.365 43.806a39.063 39.063 0 0 1-6.137-6.058m-7.333-7.201c-1.46-1.532-4.98-5.088-6.401-7.726c-3.6-6.683-.816-9.092.064-10.327a70.862 70.862 0 0 1 5.227-5.227c-.065 5.856 4.87 12.394 7.332 15.626m7.229 7.215c2.031 2.266 6.952 6.55 8.449 7.315"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.612 11.668c-4.883-1.314-14.025 8.543-16.495 11.225m-6.222 7.654c-2.542 2.967-6.465 7.229-9.385 9.337m35.502-20.437c-2.679-.115-11.09 7.397-12.666 10.661m-6.118 7.64c-1.861 2.41-4.006 5.171-5.956 7.426"/>`),
		g.Group(children),
	)
}