package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NailPolishOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNailPolishOne0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="32" height="24" x="8" y="20" fill="#fff" stroke="#fff" rx="2"/><path fill="#fff" stroke="#fff" d="M17 4h14v16H17z"/><path fill="#000" stroke="#000" d="M22 32h4l1 5h-6l1-5Z"/><path stroke="#000" d="M24 20v12"/><path stroke="#fff" d="M31 20H17"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNailPolishOne0)"/>`),
		g.Group(children),
	)
}