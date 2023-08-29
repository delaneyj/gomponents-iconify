package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UbisoftConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.503 30.132a7.511 7.511 0 0 0-12.836-7.34a7.732 7.732 0 0 0-1.774 4.968a8.037 8.037 0 0 0 1.798 4.97a8.637 8.637 0 0 0 4.482 2.852a9.676 9.676 0 0 0 5.328-.156a11.004 11.004 0 0 0 4.54-2.826a12.766 12.766 0 0 0 2.824-4.557a14.314 14.314 0 0 0-.45-10.68a14.978 14.978 0 0 0-7.68-7.468a15.71 15.71 0 0 0-10.722-.53a16.52 16.52 0 0 0-8.643 6.404a17.579 17.579 0 0 0-2.974 8.407A19.085 19.085 0 0 0 4.776 33a20.21 20.21 0 0 0 7.686 9.234a20.626 20.626 0 0 0 11.571 3.26a21.06 21.06 0 0 0 11.437-3.726a21.514 21.514 0 0 0 7.557-9.368a21.609 21.609 0 0 0-9.39-27.156a21.609 21.609 0 0 0-28.062 6.18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.503 30.132A10.487 10.487 0 0 0 12.64 19.736a10.625 10.625 0 0 0-3.84 7.473a13.41 13.41 0 0 0 .351 3.778a18.508 18.508 0 0 0 16.706 14.385"/>`),
		g.Group(children),
	)
}