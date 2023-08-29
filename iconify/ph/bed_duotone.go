package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M240 112v56H104V80h104a32 32 0 0 1 32 32Z" opacity=".2"/><path d="M208 72H24V48a8 8 0 0 0-16 0v160a8 8 0 0 0 16 0v-32h208v32a8 8 0 0 0 16 0v-96a40 40 0 0 0-40-40ZM24 88h72v72H24Zm88 72V88h96a24 24 0 0 1 24 24v48Z"/></g>`),
		g.Group(children),
	)
}