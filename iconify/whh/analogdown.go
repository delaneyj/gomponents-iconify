package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analogdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m659 843l30-50q10-12 10.5-33T689 722t-32-17l-418-1q-20 0-31.5 17t-11 38t10.5 34l31 50Q130 786 65 680T0 448q0-91 35.5-174T131 131t143-95.5T448 0t174 35.5T765 131t95.5 143T896 448q0 126-64.5 231.5T659 843zm-60-75q9 13 9 31t-9 30l-129 182q-9 13-21.5 13t-21.5-13L297 829q-8-13-8-30.5t8-30.5h302z"/>`),
		g.Group(children),
	)
}