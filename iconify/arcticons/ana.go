package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.686 12.366l25.663-3.797c3.127-.463 5.867 2.117 5.593 5.267l-2.025 23.257a3.78 3.78 0 0 1-3.99 3.446l-19.903-1.185a4.472 4.472 0 0 1-4.085-3.432L6.478 17.099a3.88 3.88 0 0 1 3.208-4.733Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.65 12.68l.011-.011a5.797 5.797 0 0 1 3.984-3.232L31.426 5.14c2.846-.655 5.627 1.031 6.454 3.727m2.137 15.593v.01l1.61 12.167c.473 3.597-2.523 6.69-6.131 6.335l-24.558-2.47c-2.792-.279-4.833-2.749-4.596-5.551l1.18-13.434"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.308 17.267v14.246a1.84 1.84 0 0 0 1.84 1.841h11.966V19.108H19.15a1.84 1.84 0 0 1-1.84-1.841h0a1.84 1.84 0 0 1 1.84-1.841h11.965M20.329 26.021h8.224m-8.224 3.631h8.224m-8.224-7.263h8.224"/>`),
		g.Group(children),
	)
}