package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskLamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M6 0a2 2 0 0 0-2 2v3.063C1.669 6.123 0 8.886 0 11.5c0 0 1.275 1.5 6 1.5s6-1.5 6-1.5c0-2.613-1.669-5.378-4-6.438v-.25l6.406-.625c.365.49.936.813 1.594.813c.138 0 .276-.005.406-.031l5.688 6.406A2.01 2.01 0 0 0 22 12c0 .316.088.61.219.875L16.406 23H12a2 2 0 0 0-2 2v1h14v-1a2 2 0 0 0-2-2h-3.281l5.156-9c.042.003.083 0 .125 0a2 2 0 0 0 0-4c-.138 0-.276.005-.406.031l-5.688-6.406A2.01 2.01 0 0 0 18 3a2 2 0 0 0-3.813-.844L8 2.781V2a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}