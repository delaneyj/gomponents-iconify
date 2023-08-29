package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MercadoPago(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.976 30.646a5.276 5.276 0 0 1-2.172-2.034a54.661 54.661 0 0 1-8.748 1.017c-3.7 0-6.687-.176-5.467-3.624s4.458-10.556 5.593-11.986s2.686-3.24 3.449-3.155c.946.106 2.715 1.284 2.516 2.034c-.189.715-1.128 2.277-2.747 1.165"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.864 14.84a8.134 8.134 0 0 1 1.293-1.641m1.783.207c.567-.144 1.64.582 1.034 1.328a4.778 4.778 0 0 1-2.817 1.324c-.625.056-2.797-.02-2.797-.02c-.925 1.597-.715 4.037-.8 6.182a9.346 9.346 0 0 1-.882 3.659c3.7-1.725 10.009-3.029 13.822-2.398M7.024 30.646a5.276 5.276 0 0 0 2.172-2.034a54.661 54.661 0 0 0 8.748 1.017c3.7 0 6.687-.176 5.467-3.624s-4.458-10.556-5.593-11.986s-2.686-3.24-3.449-3.155c-.946.106-2.715 1.284-2.516 2.034c.189.715 1.128 2.277 2.747 1.165"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.136 14.84a8.134 8.134 0 0 0-1.292-1.641m-1.784.207c-.567-.144-1.64.582-1.034 1.328a4.778 4.778 0 0 0 2.818 1.324c.624.056 2.796-.02 2.796-.02c.925 1.597.715 4.037.8 6.182a9.346 9.346 0 0 0 .882 3.659C14.622 24.153 8.313 22.85 4.5 23.48m27.681-11.738a27.566 27.566 0 0 0-16.364 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.917 13.99C6.976 16.46 4.5 20.03 4.5 24c0 7.456 8.73 13.5 19.5 13.5S43.5 31.456 43.5 24c0-3.97-2.476-7.54-6.417-10.01"/>`),
		g.Group(children),
	)
}