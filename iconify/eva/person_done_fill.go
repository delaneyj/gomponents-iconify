package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonDoneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPersonDoneFill0"><g id="evaPersonDoneFill1"><path id="evaPersonDoneFill2" fill="currentColor" d="M21.66 4.25a1 1 0 0 0-1.41.09l-1.87 2.15l-.63-.71a1 1 0 0 0-1.5 1.33l1.39 1.56a1 1 0 0 0 .75.33a1 1 0 0 0 .74-.34l2.61-3a1 1 0 0 0-.08-1.41ZM10 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm6 10a1 1 0 0 0 1-1a7 7 0 0 0-14 0a1 1 0 0 0 1 1"/></g></g>`),
		g.Group(children),
	)
}