package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlatShoe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#ffa7c0" stroke="#ffa7c0" stroke-miterlimit="10" stroke-width="2" d="M27.294 41.682c.043 1.86 5.834 2.564 8.711 2.49c28.007-.724 30.06-2.312 29.389-8.65c-.369-3.478-9.119-6.302-17.115-5.497c-.73.074-1.352 1.877-2.082 2.003c-10.767 1.86-27.834 2.56-35.93-6.075c-1.322-.198-4.887 9.17-3.542 15.351Z"/><path fill="#3f3f3f" d="M6.427 40.77h13.912v3.417H6.427z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.427 40.77c-1.693-8.08 2.518-15.016 3.84-14.817c8.096 8.635 25.032 7.974 35.8 6.114c.843-.146 1.545-2.101 2.387-2.177c7.873-.709 15.016 1.618 16.621 4.748c3.036 5.916-.53 10.938-36.626 9.27a38.643 38.643 0 0 1-3.7-1.43"/><path d="M6.427 40.77h13.912v3.417H6.427z"/></g>`),
		g.Group(children),
	)
}