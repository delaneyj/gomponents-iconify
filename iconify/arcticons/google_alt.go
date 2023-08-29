package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.681 34.887c-1.915 1.29-4.358 2.072-7.251 2.072c-5.59 0-10.34-3.773-12.04-8.854v-.02c-.43-1.29-.684-2.658-.684-4.085s.254-2.795.684-4.085c1.7-5.082 6.45-8.854 12.04-8.854c3.166 0 5.98 1.094 8.229 3.205l6.156-6.156C35.082 4.63 30.235 2.5 24.43 2.5c-8.405 0-15.656 4.828-19.194 11.864C3.77 17.257 2.93 20.521 2.93 24s.84 6.743 2.306 9.636v.02C8.774 40.671 16.026 45.5 24.43 45.5c5.805 0 10.672-1.915 14.229-5.2c4.066-3.752 6.41-9.264 6.41-15.811c0-1.525-.136-2.99-.39-4.398h-20.25v8.326h11.572c-.508 2.678-2.033 4.945-4.32 6.47h0Z"/>`),
		g.Group(children),
	)
}