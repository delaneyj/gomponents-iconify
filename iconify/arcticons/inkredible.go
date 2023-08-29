package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inkredible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 36.927c5.402-2.741 5.402-2.786 7.865-3.19c-1.124 1.491-2.405 2.435-1.888 5.347c3.052.634 5.968-1.739 8.629-2.633c-.8 1.034-2.094 3.168-.517 3.595a5.672 5.672 0 0 0 2.67-.462m4.8-3.705l-3.371 2.43l1.093-3.954"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.728 22.343l-9.67 13.536l-2.277-1.524L32.51 20.29m0 0l8.801-12.35c.922.026 1.76.548 2.189 1.364l-7.772 13.04l-3.219-2.055Zm10.299-9.693c.377 3.184-2.869 8.258-4.217 10.806"/>`),
		g.Group(children),
	)
}