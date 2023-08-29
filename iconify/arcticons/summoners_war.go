package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SummonersWar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.087 22.547a10.64 10.64 0 0 0-5.142 3.89c-1.52 2.19-5.544 7.78-5.544 9.97s6.126 8.093 6.126 8.093s-8.54-3.175-9.39-6.125a36.397 36.397 0 0 1-.894-10.82c.268-2.594 2.19-9.881 3.22-11.446S21.403 4.573 24.086 3.5c2.638.85 6.304 5.052 8.54 7.87s6.438 13.457 6.215 15.872s-1.431 12.16-2.191 13.636a29.178 29.178 0 0 0-3.13-6.975c-.537-.938-5.187-8.316-5.187-8.316s-3.13-2.682-4.247-3.04Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.915 32.583c1.245 4.811 6.671 8.951 8.867 8.951c1.756 0 8.687-3.575 9-8.844m-13.398-6.845c.218.65-.44 3.92-1.31 4.579c0 0 6.725-.952 7.49-7.015c-.057 2.851-.975 5.909-2.34 7.015c2.283-.47 5.477-4.255 5.477-4.255"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.52 39.245a6.275 6.275 0 0 1 2.728 0M24 12.617l-3.402 2.352m1.851.486l-2.384 1.82l2.234 2.807l-3.943 2.415"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 15.973l-1.63 1.615L24 19.987m0-7.37l3.402 2.352m-1.851.486l2.384 1.82l-2.234 2.807l3.943 2.415"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 15.973l1.63 1.615L24 19.987M15.58 31.46a5.92 5.92 0 0 0 5.003 1.809"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.582 33.269a5.678 5.678 0 0 0-3.767-3.734m14.445.709a5.604 5.604 0 0 0-4.64 2.587a4.359 4.359 0 0 0 5.447-1.293m-2.265.237l-.51.204m-11.665-.519l.588.297m8.292-1.333a14.374 14.374 0 0 1 3.76-1.763m-8.891 2.463a14.507 14.507 0 0 0-1.368-1.231m-1.125-.903a63.25 63.25 0 0 1-1.227-.697m5.616 7.757a1.397 1.397 0 0 0 1.087 0"/>`),
		g.Group(children),
	)
}