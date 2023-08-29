package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardIndexDividers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19.5 2A1.5 1.5 0 0 0 18 3.5V4H5a4 4 0 0 0-4 4v19a4 4 0 0 0 4 4h22a4 4 0 0 0 4-4V8a4 4 0 0 0-4-4h-2v-.5A1.5 1.5 0 0 0 23.5 2h-4ZM29 14.535A3.982 3.982 0 0 0 27 14H14v-.5a1.5 1.5 0 0 0-1.5-1.5h-4A1.5 1.5 0 0 0 7 13.5v.5H5c-.729 0-1.412.195-2 .535V13a2 2 0 0 1 2-2h22a2 2 0 0 1 2 2v1.535ZM3 18a2 2 0 0 1 2-2h22a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-9Zm26-8.465A3.981 3.981 0 0 0 27 9h-7v-.5A1.5 1.5 0 0 0 18.5 7h-4A1.5 1.5 0 0 0 13 8.5V9H5c-.729 0-1.412.195-2 .535V8a2 2 0 0 1 2-2h22a2 2 0 0 1 2 2v1.535Z"/>`),
		g.Group(children),
	)
}