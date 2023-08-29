package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HiosLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-31 13.615v10.353m6.859-10.353v10.353M9.5 24.272h4.137"/><circle cx="19.333" cy="19.438" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.333 22.609v6.859m12.513-1.135c.635.827 1.431 1.135 2.539 1.135h1.532a2.583 2.583 0 0 0 2.583-2.583v-.011a2.583 2.583 0 0 0-2.583-2.583h-1.69a2.585 2.585 0 0 1-2.586-2.585h0a2.591 2.591 0 0 1 2.591-2.591h1.525c1.107 0 1.903.308 2.538 1.134"/><rect width="6.859" height="10.353" x="22.234" y="19.115" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.429" ry="3.429"/>`),
		g.Group(children),
	)
}