package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAutofitHeightDottedTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.712 5.233L14.251 3.72L12.79 5.233a.75.75 0 0 1-1.08-1.042l1.821-1.886a1 1 0 0 1 1.44 0l1.82 1.886a.75.75 0 0 1-1.079 1.042Zm0 9.534l-1.461 1.513l-1.461-1.513a.75.75 0 0 0-1.08 1.042l1.821 1.886a1 1 0 0 0 1.44 0l1.82-1.886a.75.75 0 0 0-1.079-1.042ZM14.252 12a.75.75 0 0 1 .75.75v.5a.75.75 0 1 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Zm.75-1.75a.75.75 0 1 1-1.5 0v-.5a.75.75 0 0 1 1.5 0v.5ZM14.253 8a.75.75 0 0 0 .75-.75v-.5a.75.75 0 1 0-1.5 0v.5c0 .414.336.75.75.75ZM5 3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h4.25a.75.75 0 0 0 0-1.5H5a.5.5 0 0 1-.5-.5V5a.5.5 0 0 1 .5-.5h4.25a.75.75 0 0 0 0-1.5H5Z"/>`),
		g.Group(children),
	)
}