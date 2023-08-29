package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandHeartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.335 11.502h2.17a4.5 4.5 0 0 1 4.5 4.5H9.004v1h8v-1a5.578 5.578 0 0 0-.885-3h2.886a5 5 0 0 1 4.516 2.852c-2.365 3.12-6.194 5.149-10.516 5.149c-2.761 0-5.1-.591-7-1.625v-9.304a6.966 6.966 0 0 1 3.33 1.428Zm-5.33-2.5a1 1 0 0 1 .993.884l.007.116v9a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1h2Zm9.646-5.424l.354.354l.353-.354a2.5 2.5 0 0 1 3.536 3.535l-3.89 3.89l-3.888-3.89a2.5 2.5 0 1 1 3.535-3.535Z"/>`),
		g.Group(children),
	)
}