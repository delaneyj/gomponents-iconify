package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeneficialStateBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.538 21.358c0 9.945-1.806 16.225-13.507 22.204c-11.702-5.979-13.507-12.259-13.507-22.204M18.7 33.21h-6.18m16.28 0h6.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.81 21.36h8.73s1.77-.24 1.77-3.58c0-5.97-6.03-6.86-12.84-6.99c.16-1.1.6-2.44 1.72-3.52l-2.43-2.71S22.64 7 22.05 10.78c-7 .1-13.3.89-13.3 7c0 3.34 1.77 3.58 1.77 3.58h8.18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.3 27.276a4 4 0 0 1 0 8h-6.6v-16h6.6a4 4 0 0 1 0 8h0Zm0 0h-6.6"/>`),
		g.Group(children),
	)
}