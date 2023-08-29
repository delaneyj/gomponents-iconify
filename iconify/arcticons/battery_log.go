package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryLog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.076 32.664v7.939a1.897 1.897 0 0 1-1.898 1.897H9.471a1.897 1.897 0 0 1-1.897-1.897v-29.41A1.897 1.897 0 0 1 9.47 9.295h4.089V6.449a.949.949 0 0 1 .949-.949h6.328a.949.949 0 0 1 .949.949v2.846h4.392a1.897 1.897 0 0 1 1.898 1.897v4.144M7.574 35.859h20.502M13.56 9.295h8.226"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.93 15.336h20.29a1 1 0 0 1 1 1v12.85h0h-22.29h0v-12.85a1 1 0 0 1 1-1Zm-2.205 13.851h24.701v2.477a1 1 0 0 1-1 1h-22.7a1 1 0 0 1-1-1v-2.477h0Zm11.251 1.738h2.2"/>`),
		g.Group(children),
	)
}