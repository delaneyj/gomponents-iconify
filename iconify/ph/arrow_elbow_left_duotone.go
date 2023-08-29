package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m96 80l-72 72V80Z" opacity=".2"/><path d="M237.66 90.34a8 8 0 0 0-11.32 0L136 180.69L71.31 116l30.35-30.34A8 8 0 0 0 96 72H24a8 8 0 0 0-8 8v72a8 8 0 0 0 13.66 5.66L60 127.31l70.34 70.35a8 8 0 0 0 11.32 0l96-96a8 8 0 0 0 0-11.32ZM32 132.69V88h44.69Z"/></g>`),
		g.Group(children),
	)
}