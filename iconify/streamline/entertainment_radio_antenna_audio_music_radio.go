package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentRadioAntennaAudioMusicRadio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9.5" x=".5" y="4" rx="1"/><path d="M2 4L13.5.5"/><circle cx="4.5" cy="10" r="1.5"/><path d="M9.5 8.75H11m-1.5 2.5H11M.5 6.5h13"/></g>`),
		g.Group(children),
	)
}