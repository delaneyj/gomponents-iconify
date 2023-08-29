package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EcologyScienceAlienExtraterristerialLifeFormSpaceUniverseHead(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 5.5c0 3.59-2.95 8-6.5 8S.5 9.09.5 5.5S3.41.5 7 .5s6.5 1.41 6.5 5Z"/><path d="M2.75 4.75a3.12 3.12 0 0 0 .49 2.44a3.12 3.12 0 0 0 2.44.49a3.12 3.12 0 0 0-.49-2.44a3.12 3.12 0 0 0-2.44-.49Zm8.5 0a3.12 3.12 0 0 1-.49 2.44a3.12 3.12 0 0 1-2.44.49a3.12 3.12 0 0 1 .49-2.44a3.12 3.12 0 0 1 2.44-.49Z"/></g>`),
		g.Group(children),
	)
}