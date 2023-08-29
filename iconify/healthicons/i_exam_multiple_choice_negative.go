package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IExamMultipleChoiceNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsIExamMultipleChoiceNegative0)"><path d="M21 14a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2h-8Zm-1 5a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1Zm1 8a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2h-8Zm-1 5a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1Z"/><path fill-rule="evenodd" d="M16 26h-5a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1Zm-4 5v-3h3v3h-3Z" clip-rule="evenodd"/><path d="M17.707 14.293a1 1 0 0 1 0 1.414L13 20.414l-2.707-2.707a1 1 0 1 1 1.414-1.414L13 17.586l3.293-3.293a1 1 0 0 1 1.414 0Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm39 13a3 3 0 0 0-3 3v2h6v-2a3 3 0 0 0-3-3Zm-3 7v16.5l3 4.5l3-4.5V20h-6ZM6 39V9a3 3 0 0 1 3-3h22a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsIExamMultipleChoiceNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}