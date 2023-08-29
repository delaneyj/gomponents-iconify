package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill-rule="evenodd" d="M18 7.5H6a.5.5 0 0 0-.5.5c0 4.668 2.874 8.5 6.5 8.5s6.5-3.832 6.5-8.5a.5.5 0 0 0-.5-.5Zm-.512 1c-.19 3.932-2.608 7-5.488 7c-2.88 0-5.298-3.068-5.488-7h10.976Z" clip-rule="evenodd"/><path d="m16.862 13.329l.276-.961c.303.086.63.132.97.132c1.352 0 2.392-.72 2.392-1.5s-1.04-1.5-2.393-1.5v-1C19.95 8.5 21.5 9.572 21.5 11s-1.55 2.5-3.393 2.5c-.431 0-.852-.058-1.245-.171ZM6.5 17.75h11a.75.75 0 0 1 0 1.5h-11a.75.75 0 0 1 0-1.5Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}