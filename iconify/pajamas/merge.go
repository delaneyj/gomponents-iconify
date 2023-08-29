package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Merge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-.044 2.31a2.5 2.5 0 1 0-1.706.076v4.228a2.501 2.501 0 1 0 1.5 0V8.373a5.735 5.735 0 0 0 3.86 1.864a2.501 2.501 0 1 0 .01-1.504a4.254 4.254 0 0 1-3.664-2.922ZM11.5 10.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-6 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}