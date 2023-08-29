package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseVacancyButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52 2H12C6.477 2 2 6.476 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.524-4.477-10-10-10zM20.073 21.859v4.604h-4.53v-8.891h14.018V14h4.725v3.572h14.367v8.215h-4.764v-3.928h-5.461v6.271c0 1.07.232 1.23 1.434 1.23h3.447c.967 0 1.16-.396 1.316-2.898c.852.674 2.748 1.35 3.949 1.627c-.465 4.168-1.703 5.28-4.801 5.28h-4.607c-4.338 0-5.268-1.27-5.268-5.2v-6.311h-3.757c-.697 6.588-2.827 10.756-12.663 12.94c-.387-1.111-1.549-3.016-2.479-3.971c8.442-1.508 9.797-4.286 10.417-8.969h-5.343zM49 50H15.117v-4.287H29.56v-5.912H19.183v-4.248h26.526v4.248H34.285v5.912H49V50z"/>`),
		g.Group(children),
	)
}