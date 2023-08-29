package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 0a2 2 0 0 0-2 2v10.06l1.28-1.28l1.53-1.53H4V11a2 2 0 0 0 2 2h7l1.5 1.5L16 16V6a2 2 0 0 0-2-2h-2V2a2 2 0 0 0-2-2H2Zm8.5 4V2a.5.5 0 0 0-.5-.5H2a.5.5 0 0 0-.5.5v6.44l.47-.47l.22-.22H4V6a2 2 0 0 1 2-2h4.5Zm3.56 7.94l.44.439V6a.5.5 0 0 0-.5-.5H6a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h7.621l.44.44Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}