package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassBigBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.554 21.529c1.797 1.221 4.943-.038 11.236-2.554c1.342-.537 2.013-.806 2.54-1.267c.134-.118.26-.244.378-.378c.461-.527.73-1.198 1.267-2.54c2.515-6.293 3.775-9.44 2.554-11.236a4.085 4.085 0 0 0-1.083-1.083c-1.797-1.221-4.944.037-11.236 2.554c-1.342.537-2.013.806-2.54 1.267c-.134.118-.26.244-.378.378c-.461.527-.73 1.198-1.267 2.54c-2.517 6.292-3.775 9.439-2.554 11.236c.29.426.657.793 1.083 1.083ZM8.25 12a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Zm1.5 0a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}