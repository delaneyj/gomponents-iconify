package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForbiddenBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="m3.757 17.394l1.425 1.424l1.424 1.425c.868.867 1.301 1.3 1.852 1.529C9.01 22 9.623 22 10.85 22h2.302c1.226 0 1.84 0 2.391-.228c.551-.229.985-.662 1.852-1.53l2.849-2.848c.867-.867 1.3-1.3 1.529-1.852c.228-.551.228-1.165.228-2.39v-2.303c0-1.226 0-1.84-.228-2.39c-.228-.552-.662-.986-1.53-1.853l-1.424-1.424l-1.424-1.425c-.867-.867-1.3-1.3-1.852-1.529C14.99 2 14.377 2 13.152 2h-2.303c-1.226 0-1.84 0-2.39.228c-.552.229-.986.662-1.853 1.53L3.757 6.605c-.867.867-1.3 1.3-1.529 1.852C2 9.01 2 9.623 2 10.848v2.303c0 1.226 0 1.84.228 2.39c.229.552.662.986 1.53 1.853Z" opacity=".5"/><path stroke="currentColor" stroke-width="1.5" d="M18.818 5.182L5.182 18.818"/></g>`),
		g.Group(children),
	)
}