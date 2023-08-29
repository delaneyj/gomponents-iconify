package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M176 96a48 48 0 1 1-48-48a48 48 0 0 1 48 48Z" opacity=".2"/><path d="M216 96a88 88 0 1 0-144 67.83V240a8 8 0 0 0 11.58 7.16L128 225l44.43 22.21a8.07 8.07 0 0 0 3.57.79a8 8 0 0 0 8-8v-76.17A87.85 87.85 0 0 0 216 96ZM56 96a72 72 0 1 1 72 72a72.08 72.08 0 0 1-72-72Zm112 131.06l-36.43-18.21a8 8 0 0 0-7.16 0L88 227.06v-52.69a87.89 87.89 0 0 0 80 0ZM128 152a56 56 0 1 0-56-56a56.06 56.06 0 0 0 56 56Zm0-96a40 40 0 1 1-40 40a40 40 0 0 1 40-40Z"/></g>`),
		g.Group(children),
	)
}