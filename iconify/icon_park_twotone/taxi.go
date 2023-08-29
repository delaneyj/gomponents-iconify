package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taxi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTaxi0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M4 35v-7.29a6 6 0 0 1 3.319-5.368l.682-.34l2.31-7.91A4 4 0 0 1 14.205 11h19.688a4 4 0 0 1 3.904 3.128l2.205 7.873l.68.34A6 6 0 0 1 44 27.709V35a3 3 0 0 1-3 3h-1.997v1a4 4 0 0 1-4 4A4 4 0 0 1 31 39v-1H17v1a4 4 0 1 1-8 0v-1H7a3 3 0 0 1-3-3Z"/><path fill="#fff" d="M13.5 33a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm21 0a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 11h12l-.991-4.956C28.887 5.434 28.399 5 27.837 5h-7.674c-.563 0-1.05.434-1.172 1.044L18 11Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15 23h18"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTaxi0)"/>`),
		g.Group(children),
	)
}