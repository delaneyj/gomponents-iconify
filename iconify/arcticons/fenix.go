package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fenix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.479 5.5c-12.895 8.908-36.526.583-36.886 19.446l-.072 12.56c-.04 7.099 16.047 7.642 16.61-5.057c4.495-.727 8.71-2.064 9.563-10.756c12.162-1.41 7.326-10.732 10.785-16.194ZM5.593 24.946l26.1-3.253M22.13 32.449L5.538 34.706"/>`),
		g.Group(children),
	)
}