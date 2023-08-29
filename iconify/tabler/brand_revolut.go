package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandRevolut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.908 6c-.091.363-.908 5-.908 5h1.228C12.818 11 14 9.832 14 8.057C14 6.808 13.182 6 11.913 6h-1z"/><path d="m15.5 13.5l1.791 4.558c.535 1.352 1.13 2.008 1.709 2.442c-1 .5-2.616.522-3.605.497c-.973 0-2.28-.24-3.106-2.6L11 15H9.5s-.465 2.243-.65 3.202c-.092.704.059 1.594.15 2.298c-1 .5-2.5.5-3.5.5c-.727 0-1.45-.248-1.5-1.5v-.311a7 7 0 0 1 .149-1.409c.75-3.577 1.366-7.17 1.847-10.78c.23-1.722 0-3.5 0-3.5c.585-.144 2.709-.602 6.787-.471a10.26 10.26 0 0 1 3.641.722c.308.148.601.326.875.531c.254.212.497.437.727.674c.3.382.545.804.727 1.253c.155.483.237.987.243 1.493c0 2.462-1.412 4.676-3.5 5.798z"/></g>`),
		g.Group(children),
	)
}