package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreadmillLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="15" cy="4" r="2" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11 16v-1.633a1.85 1.85 0 0 0-.666-1.422l-.996-.83a1.59 1.59 0 0 1-.106-2.346l1.654-1.654a1.067 1.067 0 0 0-.335-1.736a4.269 4.269 0 0 0-3.943.304L4.5 8M7 14l-.328.328c-.578.579-.868.867-1.235 1.02c-.368.152-.776.152-1.594.152H3m9.5-5.5h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M19.489 22H3.087a1.087 1.087 0 0 1-.188-2.158l16.157-2.827A2.511 2.511 0 1 1 19.489 22Z"/><path fill="currentColor" d="m19.292 8.889l-.742-.111l.742.111Zm1.585-1.664l.147.735l-.147-.735Zm1.27.51a.75.75 0 1 0-.294-1.47l.294 1.47Zm-3.405 9.876l1.291-8.61l-1.483-.223l-1.292 8.61l1.484.223Zm2.282-9.651l1.123-.225l-.294-1.47l-1.123.224l.294 1.471ZM20.034 9c.052-.352.084-.555.123-.701a.63.63 0 0 1 .046-.128l-1.085-1.035c-.227.238-.34.51-.41.776c-.066.246-.11.547-.158.866L20.033 9Zm.696-2.51c-.316.062-.614.12-.857.199a1.73 1.73 0 0 0-.755.447l1.086 1.034l.012-.007a.635.635 0 0 1 .113-.046c.145-.046.346-.087.695-.157l-.294-1.47Z"/></g>`),
		g.Group(children),
	)
}