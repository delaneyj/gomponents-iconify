package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlpinejsDark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#242938" rx="60"/><path fill="#D8DEE9" fill-rule="evenodd" d="M180.778 84L223 126.037l-42.222 42.037l-42.223-42.037L180.778 84Z" clip-rule="evenodd"/><path fill="#77C1D2" fill-rule="evenodd" d="m75.222 84l87.532 87.148H78.31L33 126.037L75.222 84Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}