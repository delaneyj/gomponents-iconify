package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorUpDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 56v40a8 8 0 0 1-8 8h-36.5a8 8 0 0 0-5.88 2.57l-91.24 98.86A8 8 0 0 1 84.5 208H32a8 8 0 0 1-8-8v-40a8 8 0 0 1 8-8h36.5a8 8 0 0 0 5.88-2.57l91.24-98.86A8 8 0 0 1 171.5 48H224a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M224 40h-52.5a16 16 0 0 0-11.75 5.15L68.5 144H32a16 16 0 0 0-16 16v40a16 16 0 0 0 16 16h52.5a16 16 0 0 0 11.75-5.15L187.5 112H224a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm0 56h-36.5a16 16 0 0 0-11.75 5.15L84.5 200H32v-40h36.5a16 16 0 0 0 11.75-5.15L171.5 56H224Zm5.66 74.34a8 8 0 0 1-11.32 11.32L208 171.31V208a8 8 0 0 1-16 0v-36.69l-10.34 10.35a8 8 0 0 1-11.32-11.32l24-24a8 8 0 0 1 11.32 0Z"/></g>`),
		g.Group(children),
	)
}