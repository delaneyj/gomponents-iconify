package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Necktie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-miterlimit="10" stroke-width="2"><path fill="#61B2E4" stroke="#61B2E4" d="M16.132 13.778a5.96 5.96 0 0 0-5.96 5.96v38.079a5.96 5.96 0 0 0 5.96 5.96H54.21a5.96 5.96 0 0 0 5.961-5.96V19.738a5.96 5.96 0 0 0-5.96-5.96h-3.164L40.696 7.964l-8.528-.227l-8.221 2.305l-4.03 3.736h-3.785z"/><path fill="#92D3F5" stroke="#92D3F5" d="M20.73 12.597s-4.64 2.957 6.49 16.722l3.988-7.03l-6.824-8.414l-3.087-3.087m27.234.58l-2.41 2.41l-7.173 8.52l4.661 7.033c11.13-13.765 6.048-16.379 6.048-16.379"/><path fill="#8967AA" d="M36.679 31.62h-3.007l-4.026 32.158h11.05zm0-6.957h-3.007l-1.504 3.478l1.504 3.479h3.007l1.504-3.479z"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M16.132 13.778a5.96 5.96 0 0 0-5.96 5.96v38.079a5.96 5.96 0 0 0 5.96 5.96H54.21a5.96 5.96 0 0 0 5.961-5.96V19.738a5.96 5.96 0 0 0-5.96-5.96"/><path d="m24.384 13.875l7.824 8.414l-4.989 8.03c-11.13-13.764-7.318-17.541-7.318-17.541s14.378-13 30.614 0c0 0 3.81 3.776-7.318 17.541l-4.99-8.03l7.915-8.511M36.679 31.62h-3.007l-4.026 32.158h11.05z"/><path d="M36.679 24.663h-3.007l-1.504 3.478l1.504 3.479h3.007l1.504-3.479z"/></g>`),
		g.Group(children),
	)
}