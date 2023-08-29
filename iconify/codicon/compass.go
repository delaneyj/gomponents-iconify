package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="m9.009 6.991l2.024 4.042L6.991 9.01L4.967 4.967L9.009 6.99Zm.426 2.444L8.481 7.52l-1.916-.955l.954 1.917l1.916.954Z"/><path fill-rule="evenodd" d="M13.98 8.5a6.002 6.002 0 0 1-5.48 5.48V13h-1v.98a6.001 6.001 0 0 1-5.482-5.518H3v-1h-.976A6.001 6.001 0 0 1 7.5 2.02V3h1v-.98a6.001 6.001 0 0 1 5.48 5.48H13v1h.98ZM8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}