package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diskusage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.38 12.19a1.88 1.88 0 0 0-1.88 1.88V36.9a1.88 1.88 0 0 0 1.88 1.88h8.78l1.78-3.08h3.79v3.08H24l1.78-3.08h7.71l-10.92-13l-1.08-1.19l14.77-9.37ZM33.49 35.7l2.6 3.08h5.53a1.88 1.88 0 0 0 1.88-1.88V24.34H23.92m12.34-12.15l4.68-2.97v15.12"/>`),
		g.Group(children),
	)
}