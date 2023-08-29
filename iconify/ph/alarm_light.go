package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 34a94 94 0 1 0 94 94a94.11 94.11 0 0 0-94-94Zm0 176a82 82 0 1 1 82-82a82.1 82.1 0 0 1-82 82ZM60.24 28.24l-32 32a6 6 0 0 1-8.48-8.48l32-32a6 6 0 0 1 8.48 8.48Zm176 32a6 6 0 0 1-8.48 0l-32-32a6 6 0 0 1 8.48-8.48l32 32a6 6 0 0 1 0 8.48ZM184 122a6 6 0 0 1 0 12h-56a6 6 0 0 1-6-6V72a6 6 0 0 1 12 0v50Z"/>`),
		g.Group(children),
	)
}