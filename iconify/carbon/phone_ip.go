package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneIp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 14h2v2h-2zm4 0h2v2h-2zm4 0h2v2h-2zm-8 4h2v2h-2zm4 0h2v2h-2zm4 0h2v2h-2zm-8 4h2v2h-2zm4 0h2v2h-2zm4 0h2v2h-2zm-8-12h10v2H16z"/><path fill="currentColor" d="M28 6H14V5a2.002 2.002 0 0 0-2-2H8a2.002 2.002 0 0 0-2 2v1H4a2.002 2.002 0 0 0-2 2v18a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V8a2.002 2.002 0 0 0-2-2ZM8 5h4v17H8Zm20 21H4V8h2v14a2.002 2.002 0 0 0 2 2h4a2.002 2.002 0 0 0 2-2V8h14Z"/>`),
		g.Group(children),
	)
}