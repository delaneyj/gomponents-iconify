package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="ipSColorCard0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24.485 32H44v12H12.5"/></defs><mask id="ipSColorCard1"><g fill="none"><path fill="#fff" d="M10 44a6 6 0 0 0 6-6V4H4v34a6 6 0 0 0 6 6Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 44a6 6 0 0 0 6-6V23.515M10 44a6 6 0 0 1-6-6V4h12v19.515M10 44h34V32H24.485M5.757 42.243a6 6 0 0 0 8.486 0L24.485 32M16 23.515L35.015 4.5l.47-.5l8.5 8.5l-19.5 19.5"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14.243 42.243L43.985 12.5l-8.5-8.5L16 23.515"/><use href="#ipSColorCard0" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><use href="#ipSColorCard0" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 44a6 6 0 0 0 6-6V4H4v34a6 6 0 0 0 6 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSColorCard1)"/>`),
		g.Group(children),
	)
}