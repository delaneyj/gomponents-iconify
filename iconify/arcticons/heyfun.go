package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heyfun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.646 5.5h0c6.704 0 12.147 5.443 12.147 12.146v12.147H5.5V17.647C5.5 10.942 10.943 5.5 17.646 5.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.086 29.793c0 7.018 5.689 12.707 12.707 12.707S42.5 36.81 42.5 29.793s-5.689-12.707-12.707-12.707"/><circle cx="10.421" cy="37.641" r="3.426" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.342" cy="16.224" r="2.222" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="22.268" cy="16.224" r="2.222" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.228" cy="27.571" r="2.222" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.589 11.604l4.859-4.858m-4.859 0l4.859 4.858"/>`),
		g.Group(children),
	)
}