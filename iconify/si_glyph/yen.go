package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.722 1.625A.982.982 0 0 0 14.943.9a.983.983 0 0 0-.356-.669a.995.995 0 0 0-1.396.137L7.962 6.064L2.772.37a.994.994 0 0 0-1.757.527a.986.986 0 0 0 .219.724l5.798 6.4v1.994H4.03c-.564 0-1.018.463-1.018 1.036c0 .571.455.933 1.018.933h3.002v3c0 .555.42.99.956.99a.99.99 0 0 0 .987-.99v-3h2.973c.563 0 1.019-.361 1.019-.933c0-.573-.455-1.036-1.019-1.036H8.975V7.956l5.747-6.331z"/>`),
		g.Group(children),
	)
}