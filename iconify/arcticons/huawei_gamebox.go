package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiGamebox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.889 12.97l-4.168 6.937h8.336Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.091 18.788C13.114 3.334 34.507 4.18 38.297 17.78l4.979 17.875c1.198 4.301-2.648 5.915-5.79 4.169c-4.134-2.297-5.594-3.133-9.008-2.574c-1.71.28-4.935 1.09-8.057.224c-3.1-.86-6.758.285-10.154 2.462c-2.867 1.836-6.617-.245-5.595-4.169Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.721 25.11h8.337l-4.168 7.497Z"/>`),
		g.Group(children),
	)
}