package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mediapart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31.2" height="39" x="8.4" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.95"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.48 9.38h9.2m-9.2 7.32h9.2m-9.2 7.31h9.2m-9.2 7.32h9.2m-9.2 7.31h9.2m4.61-29.26h9.21m-9.21 7.32h9.21m-9.21 7.31h9.21m-9.21 7.32h9.21m-9.21 7.31h9.21"/>`),
		g.Group(children),
	)
}