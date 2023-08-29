package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugReport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-1.65 0-3.025-.825T6.8 18H4v-2h2.1q-.075-.575-.088-.988T6 14H4v-2h2q0-.675.013-1.113T6.1 10H4V8h2.8q.35-.6.788-1.088T8.6 6.05L7 4.4L8.4 3l2.15 2.15q.575-.2 1.425-.2t1.425.2L15.6 3L17 4.4l-1.65 1.65q.575.375 1.038.863T17.2 8H20v2h-2.1q.075.575.088.95T18 12h2v2h-2q0 .6-.013 1.012T17.9 16H20v2h-2.8q-.8 1.35-2.175 2.175T12 21Zm-2-5h4v-2h-4v2Zm0-4h4v-2h-4v2Z"/>`),
		g.Group(children),
	)
}