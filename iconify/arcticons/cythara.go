package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cythara(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.6 4.51a.83.83 0 0 1 .32 0c.5.16 11.13 10.09 10.89 10.18c-.05 0-1.91-.69-4.13-1.58s-4.49-1.79-5-2l-1-.35v10.39c0 10.31 0 10.4-.35 11a4.53 4.53 0 0 1-3.86 2.38a5.42 5.42 0 0 1-1.56-.15a4.35 4.35 0 0 1-3-2.84c-.88-3 2.41-5.81 5.68-4.77a10.63 10.63 0 0 0 1.06.31c.06 0 .13-5 .15-11.13V4.82l.4-.2a1.66 1.66 0 0 1 .38-.11ZM14.09 28a11.55 11.55 0 0 0 11.55 11.59"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.18 28a15.46 15.46 0 0 0 15.46 15.5"/>`),
		g.Group(children),
	)
}