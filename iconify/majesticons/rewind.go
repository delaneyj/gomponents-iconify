package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2.933 10.8l6.667-5c.989-.742 2.4-.036 2.4 1.2v3l5.6-4.2c.989-.742 2.4-.036 2.4 1.2v10c0 1.236-1.411 1.942-2.4 1.2L12 14v3c0 1.236-1.411 1.942-2.4 1.2l-6.667-5c-.8-.6-.8-1.8 0-2.4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}