package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataConnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 30h-6v-2h4V4h-4V2h6v28z"/><path fill="currentColor" d="M24.91 19.59a4.51 4.51 0 0 0-2.66-3.24a4.55 4.55 0 0 0-3.5 0a4.94 4.94 0 0 0-.64.35l-2.81-2.81a4.53 4.53 0 1 0-1.41 1.41l2.81 2.81a4.48 4.48 0 0 0-.61 3.3a4.51 4.51 0 0 0 2.66 3.24a4.55 4.55 0 0 0 3.5 0a4.51 4.51 0 0 0 2.66-3.24a4.65 4.65 0 0 0 0-1.82ZM11.5 14a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5Z"/><path fill="currentColor" d="M8 30H2V2h6v2H4v24h4v2z"/>`),
		g.Group(children),
	)
}