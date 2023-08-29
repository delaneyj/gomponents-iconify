package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handwashing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHandwashing0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24h7v20H4z"/><path fill="#fff" d="M32 16c-1.5-3.5 4-10 4-10s5.5 6.5 4 10s-6.5 3.5-8 0Zm-1 26.5c-9 0-16-2.5-20-2.5V30c7 0 6.5-2.5 11-4s8 0 7.5 3s-5.5 6-5.5 6c8 0 8-2 12-5s8-2 8 1s-4 11.5-13 11.5Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHandwashing0)"/>`),
		g.Group(children),
	)
}