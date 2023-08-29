package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftUnwrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.75 4.25a.75.75 0 0 0 0 1.5h14.5a.75.75 0 0 0 0-1.5H.75Zm9 6a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5h-5.5Zm-4.56 0l-.97-.97a.75.75 0 0 1 1.06-1.06l2.25 2.25l.53.53l-.53.53l-2.25 2.25a.75.75 0 0 1-1.06-1.06l.97-.97H.75a.75.75 0 0 1 0-1.5h4.44Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}