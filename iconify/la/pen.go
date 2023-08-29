package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.906 3.969A4.097 4.097 0 0 0 21 5.188L5.187 21l-.062.313l-1.094 5.5l-.312 1.468l1.469-.312l5.5-1.094l.312-.063L26.813 11a4.075 4.075 0 0 0 0-5.813a4.097 4.097 0 0 0-2.907-1.218zm0 1.906c.504 0 1.012.23 1.5.719c.973.972.973 2.027 0 3l-.718.687l-2.97-2.969l.688-.718c.489-.489.996-.719 1.5-.719zm-3.593 2.844l2.968 2.969L11.188 23.78a6.813 6.813 0 0 0-2.97-2.968zM6.938 22.438a4.734 4.734 0 0 1 2.625 2.625l-3.282.656z"/>`),
		g.Group(children),
	)
}