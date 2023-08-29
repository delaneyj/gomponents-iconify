package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accessibility(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.434 4.187a.874.874 0 0 1-.67 1.04l-2.718.587a1 1 0 0 0-.79.977v.88a8 8 0 0 0 .507 2.8l1.624 4.348a.874.874 0 1 1-1.638.612l-1.554-4.16a.5.5 0 0 0-.937 0l-1.554 4.16a.874.874 0 0 1-1.638-.611l1.625-4.35a8 8 0 0 0 .506-2.8v-.879a1 1 0 0 0-.789-.977L2.69 5.226a.874.874 0 0 1 .37-1.709l3.822.826a4 4 0 0 0 1.69 0l3.822-.826a.874.874 0 0 1 1.04.67ZM7.684 0a1.749 1.749 0 1 1 0 3.497a1.749 1.749 0 0 1 0-3.497Z"/>`),
		g.Group(children),
	)
}