package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassBigBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M5.025 9.21c.537-1.342.806-2.013 1.267-2.54c.118-.134.244-.26.378-.378c.527-.461 1.198-.73 2.54-1.267c6.292-2.517 9.439-3.775 11.236-2.554c.426.29.793.657 1.083 1.083c.951 1.4.398 3.619-1.083 7.506a256.02 256.02 0 0 1-1.471 3.73c-.537 1.342-.806 2.013-1.267 2.54c-.118.134-.244.26-.378.378c-.527.461-1.198.73-2.54 1.267c-6.293 2.515-9.44 3.775-11.236 2.554a4.086 4.086 0 0 1-1.083-1.083c-.946-1.393-.403-3.596 1.06-7.446"/><path d="M12.5 14.959a3 3 0 1 1 2.459-2.459"/></g>`),
		g.Group(children),
	)
}