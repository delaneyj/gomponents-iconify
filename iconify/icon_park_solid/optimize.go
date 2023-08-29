package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Optimize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOptimize0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke-linejoin="round" d="m19 8l9 8l10.032-5.89L33 21l9 8l-12-1l-4.5 10L23 27l-12-1l10.508-6.35L19 8Z"/><path d="M8 42.02L23 27"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOptimize0)"/>`),
		g.Group(children),
	)
}