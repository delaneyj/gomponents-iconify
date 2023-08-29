package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alembic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M17.974 37.97c2.308 5.382 8.322 10.764 14.7 10.764s11.912-5.645 14.7-10.764Z"/><path fill="#61b2e4" d="M41.749 37.97c-2.46 4.516-6.973 9.914-12.4 10.618a17.218 17.218 0 0 0 2.217.146c6.379 0 14.129-5.645 16.917-10.764Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m49.302 32.588l7.673 5.465l2.303-2.236S46.837 16.442 32.675 16.442a16.392 16.392 0 0 0-16.627 16.146a16.392 16.392 0 0 0 16.627 16.146a16.392 16.392 0 0 0 16.627-16.146"/><path d="m12.722 55.558l5.543-10.764h28.82l5.542 10.764M32.675 44.794v10.764"/></g>`),
		g.Group(children),
	)
}