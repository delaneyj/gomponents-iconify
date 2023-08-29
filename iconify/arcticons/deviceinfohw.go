package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviceinfohw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="21.663" height="39" x="13.168" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.168 38.742h21.664m0-29.484H13.168m10.832 0v29.484M27.416 13.5h4v4h-4zm0 8.5h4v4h-4zm0 8.5h4v4h-4zm4 2h3.416m-7.416-17H24m5.416 2V22m0 4v4.5"/>`),
		g.Group(children),
	)
}