package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elektrichki(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.456 7.879h21.088a2.478 2.478 0 0 1 2.483 2.473v24.216a2.478 2.478 0 0 1-2.472 2.485h-21.1a2.478 2.478 0 0 1-2.482-2.474V10.363a2.478 2.478 0 0 1 2.471-2.484Zm2.615 32.401H31.93M14.27 43.5l3.604-6.447m12.25 0l3.604 6.447"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.01 7.879h19.979v15.464H14.01z"/><circle cx="16.071" cy="32.462" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.929" cy="32.462" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.521 4.5H31.48m-4.118 0v3.379M20.638 4.5v3.379"/>`),
		g.Group(children),
	)
}