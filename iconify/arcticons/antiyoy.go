package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Antiyoy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.27 42.5l.84-1.5h3.25l1.58-2.74l-1.48-2.56h-3.22l-1.64-2.84h-3.14l-1.54-2.66l1.6-2.77h3.15l1.66-2.88l-1.57-2.73l1.43-2.47h3.3l1.76-3.05h2.84l1.62 2.81H40l1.55-2.69h1m-10.3-.12l-1.64-2.83H27.1l-1.6-2.78l1.61-2.79l-1.35-2.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.19 19.35l-1.61-2.8h-3l-1.76-3.06l1.64-2.85h3.05m.25 11.18h-3.39l-1.68 2.92h-2.92l-1.63 2.81H13l-1.57 2.71H8l-2.5 4.43m15.31-21.2H18l-1.69 2.93h-3.4l-1.69 2.88H7.6L5.5 23m2.1-3.7l-1.77-3.08l2-2.46l-1.89-3m16.58 16.67l-1.83-2.69m1.76-14.1L20.91 8h-3l-1.82 3.15h-3.16l-1.71-3l1.5-2.66m-1.5 2.67H8.11l-2.22 2.57"/>`),
		g.Group(children),
	)
}