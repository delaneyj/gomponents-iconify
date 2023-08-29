package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tradfrimelodi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.935 43.5h14.13M24 14.808V43.5m5.047-39H18.953l-5 10.308h20.094l-5-10.308z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="28.963" cy="36.79" r="2.888"/><path d="m33 26.907l-3.224.686l2.017 8.631"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="18.658" cy="26.951" r="1.723"/><path d="m21.067 21.054l-1.924.41l1.203 5.149"/></g>`),
		g.Group(children),
	)
}