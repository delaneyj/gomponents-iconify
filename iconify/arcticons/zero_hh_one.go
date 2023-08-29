package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeroHHOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="arcticons0hH10" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 39.963H22"/><path id="arcticons0hH11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.463H22"/></defs><rect width="16.5" height="16.5" x="5.5" y="26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><use href="#arcticons0hH10" stroke-linecap="round" stroke-linejoin="round"/><rect width="16.5" height="16.5" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><use href="#arcticons0hH11" stroke-linecap="round" stroke-linejoin="round"/><rect width="16.5" height="16.5" x="5.5" y="26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><use href="#arcticons0hH10" stroke-linecap="round" stroke-linejoin="round"/><rect width="16.5" height="16.5" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><use href="#arcticons0hH11" stroke-linecap="round" stroke-linejoin="round"/><rect width="16.5" height="16.5" x="26" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26 19.463h16.5"/>`),
		g.Group(children),
	)
}