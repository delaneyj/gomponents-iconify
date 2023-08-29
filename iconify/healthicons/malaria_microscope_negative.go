package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MalariaMicroscopeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsMalariaMicroscopeNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20ZM10.745 21.309c1.18 1.527 4.627.84 7.699-1.533c3.071-2.374 4.605-5.536 3.425-7.063c-1.18-1.528-4.627-.841-7.699 1.533c-3.072 2.373-4.605 5.536-3.425 7.063Zm21.694 5.541c3.287 1.985 5.151 4.918 4.165 6.552c-.987 1.634-4.45 1.35-7.737-.635c-3.287-1.984-5.152-4.918-4.165-6.552c.987-1.634 4.45-1.35 7.737.635Zm.832-3.995c1.749-.818 1.833-4.333.187-7.852c-1.646-3.518-4.398-5.707-6.147-4.889c-1.75.819-1.833 4.334-.187 7.852c1.645 3.519 4.398 5.707 6.147 4.89ZM21.62 29.933c1.784 3.55 1.8 7.145.036 8.032c-1.765.887-4.641-1.27-6.426-4.82c-1.784-3.549-1.8-7.145-.035-8.032c1.764-.887 4.64 1.271 6.425 4.82Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsMalariaMicroscopeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}