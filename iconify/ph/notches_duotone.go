package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotchesDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 40v152H40Z" opacity=".2"/><path d="M195.06 32.61a8 8 0 0 0-8.72 1.73l-152 152A8 8 0 0 0 40 200h152a8 8 0 0 0 8-8V40a8 8 0 0 0-4.94-7.39ZM184 184H59.31L184 59.31Z"/></g>`),
		g.Group(children),
	)
}