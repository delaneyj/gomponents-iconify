package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percentage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPercentage0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="11" cy="11" r="5" fill="#fff"/><circle cx="37" cy="37" r="5" fill="#fff"/><path d="M42 6L6 42"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPercentage0)"/>`),
		g.Group(children),
	)
}