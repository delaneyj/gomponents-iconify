package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 853h171v1195H341V853h171V512q0-106 40-199t109-163T824 40t200-40q106 0 199 40t163 109t110 163t40 200v341zM683 512v341h682V512q0-70-27-132t-73-109t-108-73t-133-27q-71 0-133 27t-108 73t-73 108t-27 133zm853 1365v-853H512v853h1024z"/>`),
		g.Group(children),
	)
}