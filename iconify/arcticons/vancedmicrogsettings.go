package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vancedmicrogsettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.78 2.5A21.54 21.54 0 0 1 36 6.13l-5.63 8.41a11.39 11.39 0 1 0 3.85 14.52H24V18.94h21.5V24A21.5 21.5 0 1 1 23.24 2.51Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.4 8.6a16.37 16.37 0 0 0-5.4-.91m-12.59 5.95a16.3 16.3 0 0 0 9.13 26.3m15.69-21c1.77 0 1.85 4.77 3.45 4.77s1.85-4.77 3.45-4.77"/><circle cx="39.68" cy="26.33" r=".75" fill="currentColor"/><circle cx="22.55" cy="40.19" r=".75" fill="currentColor"/><circle cx="31.24" cy="9.35" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}