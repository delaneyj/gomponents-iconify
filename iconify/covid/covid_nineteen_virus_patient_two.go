package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CovidNineteenVirusPatientTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 20.679a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858Zm-.571-9.429h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.853 3.576h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819m-3.576-1.853v-1.142m0 .571h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M6 6.75a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm4.555 4.138A5.251 5.251 0 0 0 .75 13.5v2.25H3l.75 7.5h4.5l.323-3.232"/>`),
		g.Group(children),
	)
}