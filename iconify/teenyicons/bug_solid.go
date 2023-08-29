package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.941V3.5a3.5 3.5 0 1 1 7 0v.441l2.776-1.388l.448.894l-2.953 1.477l.043.086A6.5 6.5 0 0 1 12 7.916V8h3v1h-3v1.5c0 .625-.127 1.22-.358 1.762l2.582 1.29l-.448.895l-2.627-1.313A4.494 4.494 0 0 1 7.5 15a4.494 4.494 0 0 1-3.65-1.866l-2.626 1.313l-.448-.894l2.582-1.291A4.486 4.486 0 0 1 3 10.5V9H0V8h3v-.084a6.5 6.5 0 0 1 .686-2.906l.043-.086L.776 3.447l.448-.894L4 3.94ZM5 3.5a2.5 2.5 0 0 1 5 0V4H5v-.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}