package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sailboat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="M39 5.762V37h18.687z"/><path fill="#A57939" d="M54 50c3-3 5-8 5-8H16s-1 5 3 8h35z"/><path fill="#D0CFCE" d="M35.262 13v24H23.738z"/><path fill="none" d="M69 48.235L68 68H4V48.235"/><path fill="#92D3F5" d="M68 47v21H4V47h1c3 0 9 3 15 3s10-3 16-3s11 3 16 3s7-3 15-3h1z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10"><path stroke-width="2.036" d="M39 5.762V37h18.687z"/><path stroke-width="2" d="M54.119 49.81C57.119 46.81 59 42 59 42H16s-1.208 4.838 2.792 7.838M35.262 13v24H23.738z"/><path stroke-width="2.036" d="M39 6v36"/><path stroke-width="2" d="M5 47c3 0 9 3 15 3s10-3 16-3s11 3 16 3s7-3 15-3"/></g>`),
		g.Group(children),
	)
}