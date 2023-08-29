package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransmissionVirusHumanTransmitTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 19.491a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858Zm-.571-9.428h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.818m3.576 1.853v1.143m0-.572h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.818m-1.853 3.575l-1.142.001m.571-.001v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.818m-3.576-1.853v-1.143m0 .572h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.818M4.5 6.188a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25ZM2.25 20.812a2.25 2.25 0 0 0 4.5 0v-3.75a1.5 1.5 0 0 0 1.5-1.5v-4.5a3 3 0 0 0-3-3h-1.5a3 3 0 0 0-3 3v4.5a1.5 1.5 0 0 0 1.5 1.5v3.75Zm16-19.499l2.25 2.25l-2.25 2.25m2.25-2.25l-6-.001"/>`),
		g.Group(children),
	)
}