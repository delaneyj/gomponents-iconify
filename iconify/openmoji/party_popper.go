package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartyPopper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#F1B31C" d="M26.181 20.181L38.75 32.75l12.569 12.569l-21.53 8.961L8.26 63.24l8.96-21.529z"/><path fill="#FCEA2B" d="M40 35L25.692 20.669l-9.301 21.87l-9.301 21.87z"/><path fill="#EA5A47" d="m15.224 45.285l11.062 11.062l-5.531 2.251l-7.797-7.797z"/><path fill="#D22F27" d="m17.343 55.248l3.412 3.35l5.531-2.251l-4.801-4.801z"/><path fill="#EA5A47" d="m20.805 32.163l6.696 6.696L39.2 50.558l-6.273 2.864l-8.171-8.172l-6.606-6.605z"/><path fill="#D22F27" d="m26.537 47.031l6.39 6.391l6.273-2.864l-7.838-7.838z"/><ellipse cx="30.295" cy="14.458" fill="#8967aa" rx="2" ry="1.971"/><ellipse cx="60.295" cy="18.458" fill="#f1b31c" rx="2" ry="1.971"/><ellipse cx="57.295" cy="39.458" fill="#d22f27" rx="2" ry="1.971"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m50.663 45.639l.168.169l-21.87 9.301L7.09 64.41l9.301-21.871l9.301-21.87"/><path d="m25.801 20.778l12.461 12.46l12.401 12.401m-24.971-24.97l.109.109M46.49 7.367c.235.449.403.943.49 1.473c.451 2.747-1.447 5.414-4.239 5.957"/><path d="M42.933 14.776a4.875 4.875 0 0 0-1.51.364c-2.569 1.072-3.827 4.093-2.81 6.75m23.18 4.827a4.875 4.875 0 0 1-.543 1.454c-1.372 2.423-4.523 3.31-7.037 1.98"/><path d="M54.374 30.256a4.876 4.876 0 0 0-1.403-.663c-2.673-.778-5.549.787-6.422 3.493"/></g>`),
		g.Group(children),
	)
}