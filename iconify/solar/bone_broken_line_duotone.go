package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoneBrokenLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14.137 2.738c-.743.743-1.128 2.038-.848 3.052c.149.536.173 1.185-.22 1.579l-1.28 1.28a.14.14 0 0 0 .044.226a6.268 6.268 0 0 1 3.292 3.292a.14.14 0 0 0 .227.044l1.28-1.28c.393-.393 1.042-.369 1.578-.22c1.014.28 2.309-.105 3.052-.848A2.519 2.519 0 0 0 17.7 6.3a2.519 2.519 0 0 0-3.563-3.562ZM2.738 17.7a2.519 2.519 0 0 0 3.562 0a2.519 2.519 0 0 0 3.563 3.562c.743-.743 1.128-2.038.848-3.052c-.149-.536-.173-1.185.22-1.579l1.28-1.28a.14.14 0 0 0-.044-.226a6.268 6.268 0 0 1-3.292-3.292a.14.14 0 0 0-.227-.044l-1.28 1.28c-.393.393-1.042.369-1.578.22c-1.014-.28-2.309.105-3.052.848a2.519 2.519 0 0 0 0 3.563Z"/><path d="m17.657 17.657l1.414 1.414M6.343 6.343L4.93 4.93M15 19v1M5 9H4m15 6h1M9 5V4" opacity=".5"/></g>`),
		g.Group(children),
	)
}