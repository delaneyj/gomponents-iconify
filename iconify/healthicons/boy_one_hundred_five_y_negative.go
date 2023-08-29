package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyOneHundredFiveYNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBoy0105yNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM28 13a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm6.004 8.423a2 2 0 1 0-1.008-3.87c-3.613.94-6.326 1.36-8.988 1.349c-2.668-.01-5.389-.454-9.027-1.356a2 2 0 1 0-.962 3.883c1.793.444 3.426.796 4.981 1.044v6.402l-.985 7.877a2 2 0 0 0 3.925.733l2-8A2 2 0 0 0 24 29a2 2 0 0 0 .06.485l2 8a2 2 0 0 0 3.925-.733L29 28.875v-6.376c1.56-.247 3.2-.606 5.004-1.076Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBoy0105yNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}