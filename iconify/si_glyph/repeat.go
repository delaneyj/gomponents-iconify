package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.986 4.016H6.939L4.935 2.193a.645.645 0 0 0-.911 0v1.823H1.998a2 2 0 0 0-1.998 2v3.969a2 2 0 0 0 1.998 2h7.238l1.826 1.828a.646.646 0 0 0 .912 0v-1.828h2.012a2 2 0 0 0 1.998-2V6.016a2 2 0 0 0-1.998-2zm-2.924 4.187L9.15 10.031H1.984V5.969h2.04v1.746a.631.631 0 0 0 .911 0l1.879-1.746h7.201l.002 4.062h-2.043V8.203a.632.632 0 0 0-.912 0z"/>`),
		g.Group(children),
	)
}