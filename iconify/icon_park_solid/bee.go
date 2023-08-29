package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBee0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" d="M30 28.696C30 35 27.314 44 24 44s-6-9-6-15.304C18 24.998 20.686 22 24 22s6 2.998 6 6.696Z"/><path d="M11.478 17C13.988 17 20 19.239 20 22s-6.012 5-8.522 5C8.968 27 6 24.761 6 22s2.968-5 5.478-5Zm25.044 0C34.012 17 28 19.239 28 22s6.012 5 8.522 5C39.032 27 42 24.761 42 22s-2.968-5-5.478-5Z"/><rect width="10" height="13" x="19" y="9" fill="#fff" rx="5"/><path stroke-linecap="round" d="M28 10a6 6 0 0 1 6-6m-13 6c0-3.314-3.134-6-7-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBee0)"/>`),
		g.Group(children),
	)
}