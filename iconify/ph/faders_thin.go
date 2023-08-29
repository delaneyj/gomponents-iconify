package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FadersThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132 120v96a4 4 0 0 1-8 0v-96a4 4 0 0 1 8 0Zm68 76a4 4 0 0 0-4 4v16a4 4 0 0 0 8 0v-16a4 4 0 0 0-4-4Zm24-32h-20V40a4 4 0 0 0-8 0v124h-20a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm-168 0a4 4 0 0 0-4 4v48a4 4 0 0 0 8 0v-48a4 4 0 0 0-4-4Zm24-32H60V40a4 4 0 0 0-8 0v92H32a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm72-48h-20V40a4 4 0 0 0-8 0v44h-20a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}