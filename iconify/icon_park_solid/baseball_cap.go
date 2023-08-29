package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaseballCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBaseballCap0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M12 30c0-8.837 7.163-16 16-16v0c8.837 0 16 7.163 16 16v6H12v-6Z"/><path d="M22 36c-1-3.5-1-22 6-22s6.5 18 6 22"/><rect width="30" height="6" x="4" y="36" fill="#fff" stroke-linecap="round" stroke-linejoin="round" rx="3"/><circle cx="28" cy="10" r="4" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBaseballCap0)"/>`),
		g.Group(children),
	)
}