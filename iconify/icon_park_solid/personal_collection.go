package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalCollection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPersonalCollection0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="11" r="7" fill="#fff"/><path d="M4 41c0-8.837 8.059-16 18-16"/><path fill="#fff" d="M31.85 28C29.724 28 28 30.009 28 32.486c0 4.487 4.55 8.565 7 9.514c2.45-.949 7-5.027 7-9.514C42 30.01 40.276 28 38.15 28c-1.302 0-2.453.753-3.15 1.906C34.303 28.753 33.152 28 31.85 28Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPersonalCollection0)"/>`),
		g.Group(children),
	)
}