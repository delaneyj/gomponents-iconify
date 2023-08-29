package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feUsb0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUsb1" fill="currentColor" fill-rule="nonzero"><path id="feUsb2" d="M10 4v4h6V4h-6Zm4 6v4h2v-4h-2Zm-4-8h6a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Z"/></g></g>`),
		g.Group(children),
	)
}