package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingDoNotCloseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.856 6.524a2.887 2.887 0 1 0 0-5.774a2.887 2.887 0 0 0 0 5.774ZM4.833 9.521a1.619 1.619 0 0 1-2.133.819l-.99-.44a1.618 1.618 0 0 1 1.318-2.956l.986.439a1.619 1.619 0 0 1 .819 2.138v0Zm2.306 12.11v-5.605m3.237-3.568v9.173a1.619 1.619 0 1 1-3.237 0"/><path d="M7.139 21.631a1.619 1.619 0 1 1-3.238 0v-8.173m-2.701-.25v2.421a1.617 1.617 0 0 0 1.618 1.618H3.9M15.516 6.524a2.887 2.887 0 1 0 0-5.774a2.887 2.887 0 0 0 0 5.774Z"/><path d="M21.084 10.968a5.4 5.4 0 0 0-5.312-4.445a12.107 12.107 0 0 0-1.835.053m-.863 5.343a5.4 5.4 0 0 0-5.4-5.4a7.485 7.485 0 0 0-2.844.3a1.74 1.74 0 0 0-.82.564m14.24 15.842a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm3.535-8.535l-7.07 7.07"/></g>`),
		g.Group(children),
	)
}