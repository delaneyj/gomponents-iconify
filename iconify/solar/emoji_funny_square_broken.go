package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiFunnySquareBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M8.913 15.934c1.258.315 2.685.315 4.122-.07c1.437-.385 2.673-1.099 3.605-2.001"/><ellipse cx="14.509" cy="9.774" fill="currentColor" rx="1" ry="1.5" transform="rotate(-15 14.51 9.774)"/><ellipse cx="8.714" cy="11.328" fill="currentColor" rx="1" ry="1.5" transform="rotate(-15 8.714 11.328)"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M20.796 9.643c1.111 4.147 1.667 6.22.724 7.853c-.943 1.634-3.016 2.19-7.163 3.3c-4.147 1.111-6.22 1.667-7.853.724c-1.634-.943-2.19-3.016-3.3-7.163c-1.112-4.147-1.667-6.22-.724-7.853c.943-1.634 3.016-2.19 7.163-3.3c4.147-1.112 6.22-1.667 7.853-.724c1.009.582 1.606 1.595 2.177 3.26"/><path stroke="currentColor" stroke-width="1.5" d="m13 16l.478.974a1.5 1.5 0 1 0 2.693-1.322l-.46-.935"/></g>`),
		g.Group(children),
	)
}