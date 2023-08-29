package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneRoundedBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m10.687 6.48l-.65-1.164c-.884-1.585-3.106-1.772-4.5-.379c-.836.837-1.487 1.863-1.53 2.996c-.056 1.47.217 3.728 1.83 6.173l4.278-4.278c.922-.923 1.157-2.3.572-3.349Zm3.485 7.405l-4.278 4.279c2.445 1.612 4.704 1.885 6.173 1.83c1.133-.043 2.159-.695 2.996-1.532c1.393-1.393 1.206-3.615-.38-4.5l-1.162-.649c-1.05-.585-2.426-.35-3.349.572Z" clip-rule="evenodd"/><path d="M11.024 12.976c-2.028-2.029-.91-3.148-.91-3.148l-4.278 4.278a14.247 14.247 0 0 0 1.832 2.227a14.167 14.167 0 0 0 2.226 1.83l4.278-4.277s-1.119 1.119-3.148-.91Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}