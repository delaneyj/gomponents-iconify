package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aniki(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="15.621" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.398" ry="2.905"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5h-7.171l3.981-17.426m-.031-.007a11.023 11.023 0 1 1 6.411.01"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5h7.171l-3.98-17.419l-.001-.007"/><ellipse cx="23.966" cy="15.608" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.126" ry="6.926"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.972 12.966L26.495 9.3m-.405 4.899l3.036-2.322"/>`),
		g.Group(children),
	)
}