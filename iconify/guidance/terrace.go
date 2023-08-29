package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terrace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 14v3.5a2 2 0 0 0 2 2h2V23m17-9v3.5a2 2 0 0 1-2 2h-2V23M12 8.5V15m-7 1.5h14m-7 0V23M1.5 8.5h21v-2h-.086a5 5 0 0 1-.822-.068l-.396-.066a16 16 0 0 1-8.861-4.65l-.21-.216h-.25l-.21.216a16 16 0 0 1-8.861 4.65l-.396.066a5 5 0 0 1-.822.068H1.5v2Z"/>`),
		g.Group(children),
	)
}