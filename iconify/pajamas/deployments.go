package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deployments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 1a.75.75 0 0 0 0 1.5h6.5a3.25 3.25 0 0 1 3.25 3.25V9.5h-.75a.75.75 0 0 0 0 1.5H15V5.75A4.75 4.75 0 0 0 10.25 1h-6.5ZM13 14.25a.75.75 0 0 1-.75.75h-6.5A4.75 4.75 0 0 1 1 10.25V5h2.25a.75.75 0 0 1 0 1.5H2.5v3.75a3.25 3.25 0 0 0 3.25 3.25h6.5a.75.75 0 0 1 .75.75ZM6.22 5.22a.75.75 0 0 1 1.06 0l2.25 2.25l.53.53l-.53.53l-2.25 2.25a.75.75 0 1 1-1.06-1.06L7.94 8L6.22 6.28a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}