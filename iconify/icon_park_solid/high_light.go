package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHighLight0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M6 44V25h6v-8h24v8h6v19H6Z"/><path stroke-linecap="round" d="M17 17V8l14-4v13"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHighLight0)"/>`),
		g.Group(children),
	)
}