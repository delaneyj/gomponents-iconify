package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NailPolishOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTNailPolishOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="32" height="24" x="8" y="20" fill="#555" rx="2"/><path fill="#555" d="M17 4h14v16H17zm5 28h4l1 5h-6l1-5Z"/><path d="M24 20v12m7-12H17"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTNailPolishOne0)"/>`),
		g.Group(children),
	)
}