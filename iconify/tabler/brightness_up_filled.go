package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessUpFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 8a4 4 0 1 1-3.995 4.2L8 12l.005-.2A4 4 0 0 1 12 8zm0-6a1 1 0 0 1 .993.883L13 3v2a1 1 0 0 1-1.993.117L11 5V3a1 1 0 0 1 1-1zm5.693 2.893a1 1 0 0 1 1.497 1.32l-.083.094l-1.4 1.4a1 1 0 0 1-1.497-1.32l.083-.094l1.4-1.4zM21 11a1 1 0 0 1 .117 1.993L21 13h-2a1 1 0 0 1-.117-1.993L19 11h2zm-4.707 5.293a1 1 0 0 1 1.32-.083l.094.083l1.4 1.4a1 1 0 0 1-1.32 1.497l-.094-.083l-1.4-1.4a1 1 0 0 1 0-1.414zM12 18a1 1 0 0 1 .993.883L13 19v2a1 1 0 0 1-1.993.117L11 21v-2a1 1 0 0 1 1-1zm-5.707-1.707a1 1 0 0 1 1.497 1.32l-.083.094l-1.4 1.4a1 1 0 0 1-1.497-1.32l.083-.094l1.4-1.4zM6 11a1 1 0 0 1 .117 1.993L6 13H4a1 1 0 0 1-.117-1.993L4 11h2zM4.893 4.893a1 1 0 0 1 1.32-.083l.094.083l1.4 1.4a1 1 0 0 1-1.32 1.497l-.094-.083l-1.4-1.4a1 1 0 0 1 0-1.414z"/></g>`),
		g.Group(children),
	)
}