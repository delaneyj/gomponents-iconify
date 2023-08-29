package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foodpanda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.145 34.288a4.648 4.648 0 0 0 4.711-3.555a.538.538 0 0 0-.524-.634h-8.378a.537.537 0 0 0-.525.628a4.534 4.534 0 0 0 4.716 3.561Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.179 13.311a4.75 4.75 0 0 0 1.733-3.9a5.058 5.058 0 0 0-5.2-5.055a4.92 4.92 0 0 0-4.334 2.311A18.7 18.7 0 0 0 24 4.79a19.808 19.808 0 0 0-8.378 1.878a4.92 4.92 0 0 0-4.333-2.311a5.15 5.15 0 0 0-5.2 5.055a5.456 5.456 0 0 0 1.732 3.9A18.837 18.837 0 0 0 4.5 24.145a19.5 19.5 0 0 0 39 0a18.5 18.5 0 0 0-3.321-10.834Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.18 27.555a2.498 2.498 0 0 0 1.921-.01c.982-.417 2.366-1.168 2.366-2.1c0-.868-2.456-1.012-3.322-1.012s-3.322.144-3.322 1.011c-.094.94 1.343 1.698 2.357 2.11Zm12.232-11.788c-1.733-1.155-5.488-1.733-6.933.145a2.946 2.946 0 0 0 0 4.044a15.165 15.165 0 0 1 3.612 5.634c.433 1.733 1.155 2.31 2.455 2.31s3.9-2.022 4.333-5.345a7.255 7.255 0 0 0-3.467-6.788Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.513 19.955a1.01 1.01 0 1 1 1.01-1.01a.987.987 0 0 1-1.01 1.01ZM19.668 15.91c-1.445-1.876-5.2-1.3-6.933-.144a7.415 7.415 0 0 0-3.468 6.788c.29 3.322 2.89 5.345 4.334 5.345c1.301 0 2.022-.722 2.455-2.31c.433-1.734 1.877-3.611 3.611-5.634a2.904 2.904 0 0 0 0-4.044Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.923 19.955a1.01 1.01 0 1 1 1.01-1.01a.987.987 0 0 1-1.01 1.01Z"/>`),
		g.Group(children),
	)
}