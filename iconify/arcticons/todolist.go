package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Todolist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 12.962l11.775 7.149L42.5 5.508m-37 18.644l11.775 7.149L42.5 16.699m-37 18.643l11.775 7.15L42.5 27.889"/>`),
		g.Group(children),
	)
}