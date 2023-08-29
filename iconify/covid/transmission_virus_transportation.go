package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransmissionVirusTransportation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.25 10.514a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858Zm-.571-9.429h1.142m-.571 0v2.572m3.839-1.218l.808.808m-.404-.404l-1.819 1.818m3.576 1.853v1.143m0-.572h-2.571m1.218 3.839l-.808.808m.404-.404L19.674 9.51m-1.853 3.575h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.818M11.25 7.657V6.514m0 .571h2.571m-1.218-3.838l.808-.808m-.404.404l1.819 1.818M9.75 10.915h-1.5a1.5 1.5 0 0 0-1.5 1.5V19.2m11.212 2.215h.788a1.5 1.5 0 0 0 1.5-1.5v-4.83m-13.5-2.67h-3a3 3 0 0 0-3 3v4.5a1.5 1.5 0 0 0 1.5 1.5h2.287"/><path d="M6.375 22.915a1.875 1.875 0 1 0 0-3.75a1.875 1.875 0 0 0 0 3.75Zm9.75 0a1.875 1.875 0 1 0 0-3.75a1.875 1.875 0 0 0 0 3.75Zm-7.913-1.5h6.076M.75 17.016h6"/></g>`),
		g.Group(children),
	)
}