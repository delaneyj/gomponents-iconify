package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 12h7.714v7.714H12zm16.286 0H36v7.714h-7.714zm0 16.286H36V36h-7.714zm-16.286 0h7.714V36H12z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36 19.714l-16.286 8.572L28.286 12M12 19.714l7.714 8.572v-8.572M28.286 36l-8.572-7.714h8.572"/>`),
		g.Group(children),
	)
}