package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorPickerCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M19.324 9.717a2.5 2.5 0 0 0-3.535-3.535l-8.975 8.975l-.112.17c-.165.417-.366.827-.865 1.797c-1.255 2.436-1.455 2.935-.923 3.468c.533.533 1.032.332 3.468-.923c.97-.5 1.38-.7 1.798-.865l.17-.112l8.974-8.975Zm-.707-2.828a1.5 1.5 0 0 1 0 2.121l-8.903 8.903c-.426.174-.858.387-1.79.867c-.95.49-1.351.687-1.766.853a3.757 3.757 0 0 1-.435.15c.03-.118.08-.264.15-.435c.166-.414.364-.816.853-1.766c.48-.932.693-1.364.867-1.79l8.903-8.903a1.5 1.5 0 0 1 2.121 0Z" clip-rule="evenodd"/><path d="M12.851 6.824a.5.5 0 0 1 .707-.707l5.657 5.657a.5.5 0 1 1-.707.707L12.85 6.824Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}