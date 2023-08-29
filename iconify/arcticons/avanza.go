package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Avanza(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 26.19h9.75V41.5H4.5zm9.75-6.56H24v21.88h-9.75zM24 13.06h9.75V41.5H24zm9.75-6.56h9.75v35h-9.75z"/>`),
		g.Group(children),
	)
}