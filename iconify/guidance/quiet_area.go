package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuietArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14.5 24v-4a2.5 2.5 0 0 0-2.5-2.5h-.5V24m9-4v-1.5A2.5 2.5 0 0 0 18 16h-.5v4v-9.5H17a2.5 2.5 0 0 0-2.5 2.5v8m9 3v-4a2.5 2.5 0 0 0-2.5-2.5h-.708M.5 24v-2.5a4 4 0 0 1 4-4h4v-7h3v-.25l-.063-.1A19.008 19.008 0 0 1 8.5 0"/>`),
		g.Group(children),
	)
}