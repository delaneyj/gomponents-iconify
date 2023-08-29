package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnonymousmessengerAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.286 10.071V24c0 6.757 11.293 19.497 16.714 19.5S40.714 30.757 40.714 24V10.071C40.714 4.704 29.782 4.5 24 4.5s-16.714.204-16.714 5.571Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.357 17.964a10.775 10.775 0 0 0-3.714-1.393a10.775 10.775 0 0 0-3.714 1.393a10.776 10.776 0 0 0 3.714 1.393a10.776 10.776 0 0 0 3.714-1.393Zm9.286 0a10.775 10.775 0 0 1 3.714-1.393a10.775 10.775 0 0 1 3.714 1.393a10.776 10.776 0 0 1-3.714 1.393a10.776 10.776 0 0 1-3.714-1.393Zm-6.5 17.179v4.178L24 41.18l1.857-1.858v-4.178ZM11.929 24.929l3.714 6.964h6.5L24 30.036l1.857 1.857h6.5l3.714-6.964l-4.642 4.178h-3.715l-2.785-2.786H23.07l-2.785 2.786H16.57Zm-.465-10.215c2.912-1.477 5.209-2.02 8.357 1.672a1.428 1.428 0 0 0 1.858 0a1.42 1.42 0 0 0 0-1.857c-3.78-3.85-8.265-3.483-10.215.185Zm25.072 0c-2.912-1.477-5.209-2.02-8.357 1.672a1.428 1.428 0 0 1-1.858 0a1.42 1.42 0 0 1 0-1.857c3.78-3.85 8.265-3.483 10.215.185Z"/>`),
		g.Group(children),
	)
}