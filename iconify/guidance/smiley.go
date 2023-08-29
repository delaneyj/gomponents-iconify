package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smiley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M8.5 7v3m7-3v3M12 22.5C6.201 22.5 1.5 17.799 1.5 12S6.201 1.5 12 1.5S22.5 6.201 22.5 12S17.799 22.5 12 22.5Zm.367-9.5h-.735c-2.055 0-4.078-.516-5.882-1.5H5.5v.5a6.5 6.5 0 1 0 13 0v-.5h-.25a12.284 12.284 0 0 1-5.883 1.5Z"/>`),
		g.Group(children),
	)
}