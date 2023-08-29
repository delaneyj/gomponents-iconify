package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m230.14 70.54l-44.68-44.69a20 20 0 0 0-28.29 0L33.86 149.17A19.85 19.85 0 0 0 28 163.31V208a20 20 0 0 0 20 20h44.69a19.86 19.86 0 0 0 14.14-5.86L230.14 98.82a20 20 0 0 0 0-28.28ZM93 180l71-71l11 11l-71 71Zm-17-17l-11-11l71-71l11 11Zm-24 10l15.51 15.51L83 204H52Zm140-70l-39-39l18.34-18.34l39 39Z"/>`),
		g.Group(children),
	)
}