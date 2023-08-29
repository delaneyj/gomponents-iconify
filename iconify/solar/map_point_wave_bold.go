package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPointWaveBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22c5.523 0 10-2.014 10-4.5c0-1.266-1.163-2.41-3.035-3.229c-1.142 2.096-2.883 3.903-5.095 4.848a4.775 4.775 0 0 1-3.74 0c-2.212-.945-3.953-2.752-5.095-4.848C3.163 15.09 2 16.235 2 17.502C2 19.985 6.477 22 12 22Z"/><path fill-rule="evenodd" d="M5 8.515C5 4.917 8.134 2 12 2s7 2.917 7 6.515c0 3.57-2.234 7.735-5.72 9.225a3.277 3.277 0 0 1-2.56 0C7.234 16.25 5 12.084 5 8.515ZM12 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}