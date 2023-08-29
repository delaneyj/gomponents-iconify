package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Restricted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M724 1024H300L0 724V300L300 0h424l300 300v424zm172-671L671 128H353L128 353v318l225 225h318l225-225V353z"/>`),
		g.Group(children),
	)
}