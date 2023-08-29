package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Theaters(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 22H2V2h20v20ZM20 4h-2.5v2.5H20V4Zm0 4.5h-2.5V11H20V8.5Zm0 4.5h-2.5v2.5H20V13Zm-2.5 4.5V20H20v-2.5h-2.5Zm-2-2V13H15v-2h.5V8.5H15v-2h.5V4h-7v2.5H9v2h-.5V11H9v2h-.5v2.5H9v2h-.5V20h7v-2.5H15v-2h.5ZM6.5 4H4v2.5h2.5V4Zm0 4.5H4V11h2.5V8.5Zm0 4.5H4v2.5h2.5V13Zm0 4.5H4V20h2.5v-2.5Z"/>`),
		g.Group(children),
	)
}