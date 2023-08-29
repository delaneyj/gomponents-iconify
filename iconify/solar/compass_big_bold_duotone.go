package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassBigBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M14.79 18.975C8.497 21.49 5.35 22.75 3.554 21.529a4.086 4.086 0 0 1-1.083-1.083c-1.221-1.797.037-4.944 2.554-11.236c.537-1.342.806-2.013 1.267-2.54c.118-.134.244-.26.378-.378c.527-.461 1.198-.73 2.54-1.267c6.292-2.517 9.439-3.775 11.236-2.554c.426.29.793.657 1.083 1.083c1.221 1.797-.038 4.943-2.554 11.236c-.537 1.342-.806 2.013-1.267 2.54c-.118.134-.244.26-.378.378c-.527.461-1.198.73-2.54 1.267Z" clip-rule="evenodd" opacity=".5"/><path d="M12 8.25a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5Z"/></g>`),
		g.Group(children),
	)
}