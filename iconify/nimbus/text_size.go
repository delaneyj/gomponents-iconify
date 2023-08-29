package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.1 2.75h3.09V4h1.25V1.5H.5V4h1.25V2.75h3.1v10.5H2.7v1.25h5.55v-1.25H6.1V2.75z"/><path fill="currentColor" d="M12.4 6.75H8.06v2.5h1.25V8h1.84v5.25H9.63v1.25h4.3v-1.25H12.4V8h1.85v1.25h1.25v-2.5h-3.1z"/>`),
		g.Group(children),
	)
}