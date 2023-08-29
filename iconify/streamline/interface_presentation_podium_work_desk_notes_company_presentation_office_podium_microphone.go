package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfacePresentationPodiumWorkDeskNotesCompanyPresentationOfficePodiumMicrophone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 8.5H2l-1-5h12l-1 5H9m-2-5v-3m-2 6v7m4-7v7m-5 0h6m-7.5-10s0-3 2-3m7 3s0-3-2-3"/>`),
		g.Group(children),
	)
}