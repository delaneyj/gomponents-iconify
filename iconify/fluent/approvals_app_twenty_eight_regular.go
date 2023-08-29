package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApprovalsAppTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.28 2.218a.75.75 0 0 0-1.06 1.06l1.708 1.709c-5.773.038-10.442 4.73-10.442 10.512c0 5.806 4.707 10.513 10.513 10.513c5.716 0 10.366-4.562 10.509-10.244a.75.75 0 0 0-1.5-.038a9.013 9.013 0 1 1-9.056-9.243L12.22 8.219a.75.75 0 1 0 1.06 1.06l3-3a.75.75 0 0 0 0-1.06l-3-3.001Zm5 10.001a.75.75 0 0 1 0 1.06l-5.25 5.25a.75.75 0 0 1-1.06 0l-2.252-2.25a.75.75 0 0 1 1.06-1.061l1.722 1.72l4.72-4.72a.75.75 0 0 1 1.06.001Z"/>`),
		g.Group(children),
	)
}