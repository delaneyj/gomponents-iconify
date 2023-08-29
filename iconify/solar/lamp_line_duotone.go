package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M9 22h6"/><path stroke-linecap="round" d="M12 22v-7" opacity=".5"/><path d="M4.961 7.445c.61-2.346.914-3.519 1.7-4.294a4 4 0 0 1 .757-.585C8.365 2 9.577 2 12 2s3.635 0 4.582.566a4 4 0 0 1 .757.585c.786.775 1.09 1.948 1.7 4.294l.084.324c.827 3.189 1.241 4.783.49 5.903a2.998 2.998 0 0 1-.247.319C18.47 15 16.823 15 13.529 15h-3.058c-3.295 0-4.942 0-5.837-1.01c-.09-.1-.172-.206-.247-.318c-.752-1.12-.338-2.714.49-5.903l.084-.324Z"/><path stroke-linecap="round" d="M17.5 15v2" opacity=".5"/></g>`),
		g.Group(children),
	)
}