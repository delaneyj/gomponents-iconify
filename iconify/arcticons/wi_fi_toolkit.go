package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiToolkit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="27" x="4.49" y="13.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.5 13.5v-3a3 3 0 0 0-3-3h-9a3 3 0 0 0-3 3v3m10.96 16.98a5.393 5.393 0 0 0-6.89-.018v.019M24 32.689a2.162 2.162 0 1 0 2.163 2.162v0h0a2.162 2.162 0 0 0-2.162-2.162Zm6.748-5.946a10.603 10.603 0 0 0-13.495 0m16.737-3.924a15.676 15.676 0 0 0-20 0"/>`),
		g.Group(children),
	)
}