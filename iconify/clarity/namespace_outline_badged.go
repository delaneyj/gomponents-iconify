package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NamespaceOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M9.68 29.9L4 26.62V9.38L9.68 6.1a1 1 0 1 0-.93-1.77h-.07L2 8.23v19.54l6.68 3.86a1 1 0 0 0 1.37-.36a1 1 0 0 0-.37-1.37Z"/><path fill="currentColor" d="M26 12.34A7.68 7.68 0 0 1 23.66 10H12a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2ZM17 24h-5v-5h5Zm0-7h-5v-5h5Zm7 7h-5v-5h5Zm0-7h-5v-5h5Zm8 9.62l-6 3.47a1 1 0 1 0 1 1.73l7-4.05V12.34a7.65 7.65 0 0 1-2 .88Z"/><circle cx="30" cy="6" r="5" fill="currentColor"/>`),
		g.Group(children),
	)
}