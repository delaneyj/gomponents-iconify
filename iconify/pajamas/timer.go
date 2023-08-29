package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 1a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2v.07a6.958 6.958 0 0 1 2.812 1.058l.908-.908a.75.75 0 1 1 1.06 1.06l-.8.8A7 7 0 1 1 7 2.07V2a1 1 0 0 1-1-1Zm7.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0ZM8.75 5.75a.75.75 0 0 0-1.5 0v3.56l.22.22l2.254 2.254a.75.75 0 1 0 1.06-1.06L8.75 8.689V5.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}