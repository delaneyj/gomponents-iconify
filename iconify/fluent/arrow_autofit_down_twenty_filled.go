package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAutofitDownTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m15.712 14.767l-.71.735V2.75a.75.75 0 1 0-1.5 0v12.756l-.713-.739a.75.75 0 0 0-1.079 1.042l1.821 1.886a1 1 0 0 0 1.44 0l1.82-1.886a.75.75 0 0 0-1.079-1.042ZM5 17a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6.25a.75.75 0 0 1 0 1.5H5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h4.25a.75.75 0 0 1 0 1.5H5Z"/>`),
		g.Group(children),
	)
}