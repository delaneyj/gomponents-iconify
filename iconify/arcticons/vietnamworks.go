package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vietnamworks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.845 26.893c1.238-6.187 2.97-10.394 7.672-13.364c-8.91-1.276-11.136 6.187-7.672 13.364Zm16.829 10.395c12.374-10.89 11.879-18.066 9.899-24.006c15.344 0 6.434 24.006-9.9 24.006ZM6.113 27.14c1.485 1.98 7.672 11.137 17.076 11.137c-11.88 6.93-20.046-.742-17.076-11.136Zm7.672-15.096c7.92-5.444 13.859-7.919 19.798.495c-4.95 0-7.424 1.238-9.9 1.238s-5.939-1.733-9.898-1.733Z"/>`),
		g.Group(children),
	)
}