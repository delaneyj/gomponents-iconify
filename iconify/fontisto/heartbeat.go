package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heartbeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.702-.001H4.34a4.343 4.343 0 0 0-4.341 4.34v11.964a4.344 4.344 0 0 0 4.341 4.342h3.44v1.413h-4V24h15.493v-1.942h-4v-1.413h3.434a4.344 4.344 0 0 0 4.341-4.341V4.338A4.338 4.338 0 0 0 18.71 0h-.007zm1.449 16.312v.001c0 .8-.649 1.449-1.449 1.449H4.339c-.8 0-1.448-.648-1.448-1.448v-4.98h6.856l.622-2.03l2.88 6.54l1.305-4.544h5.594v5.014zm0-6.457h-6.674l-.507 1.774l-2.834-6.435l-1.449 4.695H2.89V4.345c0-.8.649-1.449 1.449-1.45h14.362c.8 0 1.449.649 1.449 1.449v.001z"/>`),
		g.Group(children),
	)
}