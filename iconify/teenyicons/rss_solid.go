package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 15C14 7.268 7.732 1 0 1V0c8.284 0 15 6.716 15 15h-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 13.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/><path fill="currentColor" d="M9 15a9 9 0 0 0-9-9v1a8 8 0 0 1 8 8h1Z"/>`),
		g.Group(children),
	)
}