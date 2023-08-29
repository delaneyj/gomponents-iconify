package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFrameCutLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14.5 3.003c2.794.02 4.324.163 5.328 1.168C21 5.343 21 7.228 21 11v2c0 3.772 0 5.657-1.172 6.829c-1.004 1.005-2.534 1.148-5.328 1.168m-5 0c-2.794-.02-4.324-.163-5.328-1.168C3 18.656 3 16.771 3 12.999v-2c0-3.77 0-5.656 1.172-6.828C5.176 3.166 6.706 3.023 9.5 3.003"/><path d="M17 4v16"/><path stroke-dasharray="3 3" d="M12 2v20"/><path d="M7 4v16M3.5 8.5H7m13.5 0H17m-13.5 7H7m13.5 0H17"/></g>`),
		g.Group(children),
	)
}