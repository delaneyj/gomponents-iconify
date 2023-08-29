package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTaxiCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M9.171 6a2.5 2.5 0 0 0-2.425 1.894l-1.06 4.242A2.501 2.501 0 0 0 4 14.5v3A2.5 2.5 0 0 0 6.5 20h13a2.5 2.5 0 0 0 2.5-2.5v-3a2.501 2.501 0 0 0-1.685-2.364l-1.06-4.242A2.5 2.5 0 0 0 16.828 6H9.17ZM6.5 14a1 1 0 0 0 .97-.758L8.686 8.38A.5.5 0 0 1 9.171 8h7.659a.5.5 0 0 1 .485.379l1.216 4.864a1 1 0 0 0 .97.757a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-13a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path d="M19.5 16a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm-10 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM7 18.5A1.5 1.5 0 0 1 8.5 20v1a1.5 1.5 0 0 1-3 0v-1A1.5 1.5 0 0 1 7 18.5Zm12 0a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-3 0v-1a1.5 1.5 0 0 1 1.5-1.5ZM13 4a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/><path fill-rule="evenodd" d="M17.15 10.08a2 2 0 0 0-1.955-1.58h-4.39a2 2 0 0 0-1.956 1.58l-.429 2a2 2 0 0 0 1.956 2.42h5.248a1.999 1.999 0 0 0 1.955-2.42l-.429-2Zm-1.526 2.42h-5.248l.429-2h4.39l.429 2Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTaxiCircleFilled0)"/></g>`),
		g.Group(children),
	)
}