package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwatchesPalette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.5 21h13.97a.53.53 0 0 0 .53-.53v-5.94M7.98 20.67l12.662-5.904a.53.53 0 0 0 .256-.704l-2.51-5.384a.53.53 0 0 0-.704-.256l-5.654 2.636m-2.148 7.346a3.5 3.5 0 0 1-6.762-1.812L6.736 3.1a.53.53 0 0 1 .648-.375l5.74 1.538a.53.53 0 0 1 .373.648L9.882 18.405ZM6.5 17.6h.002v.002H6.5V17.6Z"/>`),
		g.Group(children),
	)
}