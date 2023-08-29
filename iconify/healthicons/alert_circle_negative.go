package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertCircleNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsAlertCircleNegative0)"><path d="M24 11a2 2 0 0 1 2 2v14a2 2 0 1 1-4 0V13a2 2 0 0 1 2-2Zm2 24a2 2 0 1 0-4 0a2 2 0 0 0 4 0Z"/><path fill-rule="evenodd" d="M0 48h48V0H0v48ZM24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20S35.046 4 24 4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsAlertCircleNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}