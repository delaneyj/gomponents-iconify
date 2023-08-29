package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9.171 6a2.5 2.5 0 0 0-2.425 1.894l-1.06 4.242A2.501 2.501 0 0 0 4 14.5v3A2.5 2.5 0 0 0 6.5 20h13a2.5 2.5 0 0 0 2.5-2.5v-3a2.501 2.501 0 0 0-1.685-2.364l-1.06-4.242A2.5 2.5 0 0 0 16.828 6H9.17ZM6.5 14a1 1 0 0 0 .97-.758L8.686 8.38A.5.5 0 0 1 9.171 8h7.659a.5.5 0 0 1 .485.379l1.216 4.864a1 1 0 0 0 .97.757a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-13a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path d="M19.5 16a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm-10 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM7 18.5A1.5 1.5 0 0 1 8.5 20v1a1.5 1.5 0 0 1-3 0v-1A1.5 1.5 0 0 1 7 18.5Zm12 0a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-3 0v-1a1.5 1.5 0 0 1 1.5-1.5ZM13 4a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/><path fill-rule="evenodd" d="M17.15 10.08a2 2 0 0 0-1.955-1.58h-4.39a2 2 0 0 0-1.956 1.58l-.429 2a2 2 0 0 0 1.956 2.42h5.248a1.999 1.999 0 0 0 1.955-2.42l-.429-2Zm-1.526 2.42h-5.248l.429-2h4.39l.429 2Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}