package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlPad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.277 7.048c-2.344 0-3.242 1.958-3.242 4.372v.166c0 2.416.898 4.373 3.242 4.373h9.475c2.344 0 3.242-1.957 3.242-4.373v-.166c0-2.414-.898-4.372-3.242-4.372H4.277zM8 12H6v2.016H5V12H3v-1h2V8.984h1V11h2v1zm4-1h-1v-1h1v1zm2 2h-1v-1h1v1z"/><path d="M9 7.104H8V4l2.031.041L10 1.047h1V5H9v2.104z"/></g>`),
		g.Group(children),
	)
}