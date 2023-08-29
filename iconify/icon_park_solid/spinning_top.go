package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinningTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSpinningTop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M24 44v-3"/><path fill="#fff" d="M44 20L24 41L4 20h40Z"/><path d="M44 12H4v8h40v-8Zm-20 0V4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSpinningTop0)"/>`),
		g.Group(children),
	)
}