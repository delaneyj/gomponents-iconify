package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlightSmile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0ZM16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0ZM6 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-4.37 4.384a2.749 2.749 0 0 0 3.751-1.009a.75.75 0 0 0-1.299-.75a1.25 1.25 0 0 1-2.163.003a.75.75 0 0 0-1.297.753c.242.417.59.763 1.007 1.003Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}