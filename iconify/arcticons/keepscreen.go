package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keepscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.853 35.444h.119a.56.56 0 0 0 .233-.045a1.305 1.305 0 0 0 .692-1.311c-.232-.864.692-2.26 1.509-3.159a11.918 11.918 0 0 0 4.555-9.268a11.018 11.018 0 0 0-22.035-.001a11.875 11.875 0 0 0 4.542 9.27a.011.011 0 0 0 .011.01c.42.467 1.638 1.909 1.557 3.09a1.695 1.695 0 0 0 .83 1.37a.631.631 0 0 0 .165.048v5.164a.799.799 0 0 0 .796.806h.023a3.336 3.336 0 0 0 6.183 0h.025a.8.8 0 0 0 .795-.806Zm-8.647 1.051h9.467m-9.467 3.47h9.467m-9.467-1.735h9.467M24.025 4.5l-.008 4.364m-12.306.666l3.081 3.091M36.288 9.65l-3.091 3.081"/>`),
		g.Group(children),
	)
}