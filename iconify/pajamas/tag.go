package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.172 5.5a.5.5 0 0 1 .353.146l1.06-1.06l-1.06 1.06L13.88 8l-2.354 2.354a.5.5 0 0 1-.353.146H2a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 .5-.5h9.172Zm3.767 1.44l-2.353-2.354A2 2 0 0 0 11.172 4H2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h9.172a2 2 0 0 0 1.414-.586l2.353-2.353L16 8l-1.06-1.06Zm-8.189.31a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}