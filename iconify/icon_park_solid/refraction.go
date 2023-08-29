package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="ipSRefraction0" fill="#fff" d="m24 9l16.454 28.5H7.545L24 9Z"/></defs><mask id="ipSRefraction1"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><use href="#ipSRefraction0"/><use href="#ipSRefraction0"/><path d="m4 22l15.5-5m8.5-1l16-3m-14 6.5L44 21m-11.3 3L44 29"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRefraction1)"/>`),
		g.Group(children),
	)
}