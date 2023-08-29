package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M157 109h-4l-22-45l-65 8l26 55q-39 16-61 41.5t-26 51T2.5 268t7.5 38l5 14l212-61l-58-119q53-12 95.5-5t69.5 22l26 15l-72 19l-20-39q-35-5-66 0l49 100l211-62q-1-2-3-4.5t-9-10.5t-16-15t-24-16.5t-31.5-17t-40.5-14t-50-9.5t-60-1.5t-71 7.5zM72 250q-2-5-4-14t7.5-31t35.5-39l32 65z"/>`),
		g.Group(children),
	)
}