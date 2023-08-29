package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReviewCheckmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM8 0a2.5 2.5 0 0 1 2.45 2H13a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h2.55A2.5 2.5 0 0 1 8 0ZM7 5h3.5V3.5h2v11h-9v-11h2V5H7Zm3.53 3.28a.75.75 0 1 0-1.06-1.06L7.5 9.19l-.47-.47a.75.75 0 0 0-1.06 1.06l1 1a.75.75 0 0 0 1.06 0l2.5-2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}