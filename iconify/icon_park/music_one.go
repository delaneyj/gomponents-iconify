package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 6V35"/><path fill="#2F88FF" d="M10 36.04C10 33.2565 12.2565 31 15.04 31H24V36.96C24 39.7435 21.7435 42 18.96 42H15.04C12.2565 42 10 39.7435 10 36.96V36.04Z"/><path stroke-linecap="round" d="M24 14.0664L36.8834 17.1215V9.01341L24 6.00002V14.0664Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}