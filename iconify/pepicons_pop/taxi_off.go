package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.171 3a2.5 2.5 0 0 0-2.425 1.894l-1.06 4.242A2.501 2.501 0 0 0 1 11.5v3A2.5 2.5 0 0 0 3.5 17h13a2.5 2.5 0 0 0 2.5-2.5v-3a2.501 2.501 0 0 0-1.685-2.364l-1.06-4.242A2.5 2.5 0 0 0 13.828 3H6.17ZM3.5 11a1 1 0 0 0 .97-.758L5.686 5.38A.5.5 0 0 1 6.171 5h7.659a.5.5 0 0 1 .485.379l1.216 4.864a1 1 0 0 0 .97.757a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-13a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path d="M16.5 13a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm-10 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM4 15.5A1.5 1.5 0 0 1 5.5 17v1a1.5 1.5 0 0 1-3 0v-1A1.5 1.5 0 0 1 4 15.5Zm12 0a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-3 0v-1a1.5 1.5 0 0 1 1.5-1.5ZM10 1a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/><path fill-rule="evenodd" d="M14.15 7.08a2 2 0 0 0-1.955-1.58h-4.39a2 2 0 0 0-1.956 1.58l-.429 2a2 2 0 0 0 1.956 2.42h5.248a1.999 1.999 0 0 0 1.955-2.42l-.429-2ZM12.624 9.5H7.376l.429-2h4.39l.429 2Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}