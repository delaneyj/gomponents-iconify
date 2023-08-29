package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeLightroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0v512h512V0H0zm261.56 363.985H128v-256.93h41.89v218.298h98.47l-6.8 38.632zm131.284-154.53c-14.662-.699-39.099 1.163-47.244 4.654v149.644h-42.124l.233-150.808c.04-7.74-.167-10.746-1.295-24.754c16.785-6.662 51.557-16.904 90.43-16.904v38.168z"/>`),
		g.Group(children),
	)
}