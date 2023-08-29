package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleaningNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCleaningNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM21.479 28.177l4.81-2.557a2 2 0 0 1 2.705.827l.648 1.217l-8.343 4.436l-.647-1.218a2 2 0 0 1 .827-2.705Zm9.281-2.67l.648 1.22l6.869 10.054l-8.737 4.645l-3.027-5.983l-1.784.903l3.045 6.02l-3.746 1.991l-4.495-11.318l-.647-1.218a4 4 0 0 1 1.654-5.41l1.522-.809L11.33 5.416a1 1 0 0 1 1.766-.939l10.733 20.186l1.522-.81a4 4 0 0 1 5.41 1.655Zm-14.4 9.724a1 1 0 0 1 1.28 1.537l-.016.013c-.56.45-1.197.813-1.844 1.122A11.15 11.15 0 0 1 11 39a1 1 0 1 1 0-2a9.148 9.148 0 0 0 3.92-.903a8.83 8.83 0 0 0 1.094-.617a6.033 6.033 0 0 0 .336-.24l.01-.009Zm3.408 5.129a1 1 0 0 0-1.408-.129l-.01.008l-.062.048c-.058.044-.15.11-.274.193a8.83 8.83 0 0 1-1.094.617A9.148 9.148 0 0 1 13 42a1 1 0 1 0 0 2a11.15 11.15 0 0 0 4.78-1.097c.647-.309 1.283-.671 1.844-1.122l.016-.013a1 1 0 0 0 .128-1.408Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCleaningNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}