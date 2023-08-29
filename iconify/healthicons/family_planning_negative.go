package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FamilyPlanningNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsFamilyPlanningNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0Zm-7 42a1 1 0 1 0 0-2h-1V22c0-2.22-1.207-4.16-3-5.197V16a3 3 0 1 0-6 0v.803A5.998 5.998 0 0 0 28 22v18h-1a1 1 0 1 0 0 2h14ZM13 27.778c0-.842.658-1.547 1.542-1.731l-.93-1.882L6.48 25.97L6 24.03l7.133-1.806a1.967 1.967 0 0 1 2.246 1.046L16 24.528l.621-1.258a1.967 1.967 0 0 1 2.246-1.046L26 24.03l-.48 1.94l-7.132-1.805l-.93 1.882c.884.184 1.542.89 1.542 1.73v10.445c0 .982-.895 1.778-2 1.778v2h-2v-2c-1.105 0-2-.796-2-1.778V27.778ZM8 6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H8Zm7 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-3-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-3-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm7-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsFamilyPlanningNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}