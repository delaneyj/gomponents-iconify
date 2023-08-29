package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmWatsonNaturalLanguageUnderstanding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 23h5v2H6zm0-4h5v2H6z"/><path fill="currentColor" d="M13 30H4c-1.1 0-2-.9-2-2V17c0-1.1.9-2 2-2h9c1.1 0 2 .9 2 2v11c0 1.1-.9 2-2 2zM4 17v11h9V17H4zM19 2h8v2h-8zm3 4h8v2h-8zm0 4h8v2h-8zm-3 4h8v2h-8zm3 4h8v2h-8zM12 1l-1.4 1.4L13.2 5H4c-1.1 0-2 .9-2 2v5h2V7h9.2l-2.6 2.6L12 11l5-5l-5-5z"/>`),
		g.Group(children),
	)
}