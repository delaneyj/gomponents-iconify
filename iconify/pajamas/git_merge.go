package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.34 1.22a.75.75 0 0 0-1.06 0L7.53 2.97L7 3.5l.53.53l1.75 1.75a.75.75 0 1 0 1.06-1.06l-.47-.47h.63c.69 0 1.25.56 1.25 1.25v4.614a2.501 2.501 0 1 0 1.5 0V5.5a2.75 2.75 0 0 0-2.75-2.75h-.63l.47-.47a.75.75 0 0 0 0-1.06ZM13.5 12.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-9 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm1.5 0a2.5 2.5 0 1 1-3.25-2.386V5.886a2.501 2.501 0 1 1 1.5 0v4.228A2.501 2.501 0 0 1 6 12.5Zm-1.5-9a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}