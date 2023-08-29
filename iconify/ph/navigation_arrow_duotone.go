package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationArrowDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m226.35 121l-76.55 23.5a8 8 0 0 0-5.3 5.3L121 226.35a8 8 0 0 1-15.21.27l-65.28-176a8 8 0 0 1 10.12-10.16l176 65.28a8 8 0 0 1-.28 15.26Z" opacity=".2"/><path d="M229.33 98.21L53.41 33l-.16-.05a16 16 0 0 0-20.35 20.3a1 1 0 0 0 .05.16l65.26 175.92A15.77 15.77 0 0 0 113.28 240h.3a15.77 15.77 0 0 0 15-11.29l23.56-76.56l76.56-23.56a16 16 0 0 0 .62-30.38ZM224 113.3l-76.56 23.56a16 16 0 0 0-10.58 10.58L113.3 224l-.06-.17L48 48l175.82 65.22l.16.06Z"/></g>`),
		g.Group(children),
	)
}