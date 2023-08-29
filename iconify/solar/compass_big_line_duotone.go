package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassBigLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14.79 18.975C8.497 21.49 5.35 22.75 3.554 21.529a4.086 4.086 0 0 1-1.083-1.083c-1.221-1.797.037-4.944 2.554-11.236c.537-1.342.806-2.013 1.267-2.54c.118-.134.244-.26.378-.378c.527-.461 1.198-.73 2.54-1.267c6.292-2.517 9.439-3.775 11.236-2.554c.426.29.793.657 1.083 1.083c1.221 1.797-.038 4.943-2.554 11.236c-.537 1.342-.806 2.013-1.267 2.54c-.118.134-.244.26-.378.378c-.527.461-1.198.73-2.54 1.267Z" opacity=".5"/><circle cx="12" cy="12" r="3"/></g>`),
		g.Group(children),
	)
}