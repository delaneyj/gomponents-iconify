package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoldMedal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGoldMedal0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 4H31l-4 10.3A15.023 15.023 0 0 1 37.27 22L44 4ZM17 4H4l6.73 18A15.023 15.023 0 0 1 21 14.3L17 4Z"/><path fill="#555" d="M39 29c0 8.284-6.716 15-15 15c-8.284 0-15-6.716-15-15c0-2.528.625-4.91 1.73-7A15.023 15.023 0 0 1 21 14.3c.97-.197 1.973-.3 3-.3s2.03.103 3 .3A15.023 15.023 0 0 1 37.27 22A14.935 14.935 0 0 1 39 29Z"/><path d="M24 35V22l-3 1m3 12h4m-4 0h-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGoldMedal0)"/>`),
		g.Group(children),
	)
}