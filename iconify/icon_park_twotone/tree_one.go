package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTreeOne0"><g fill="none" stroke="#fff" stroke-width="4"><ellipse cx="24" cy="20" fill="#555" rx="15" ry="16"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 14v22"/><path d="M30 34.669A14.154 14.154 0 0 1 24 36a14.16 14.16 0 0 1-6-1.331"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 36v8m0-22l6-6m-6 13l-6-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTreeOne0)"/>`),
		g.Group(children),
	)
}