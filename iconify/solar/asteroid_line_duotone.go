package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteroidLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 2C6.477 2 2 6.477 2 12a10.01 10.01 0 0 0 7.383 9.654A9.99 9.99 0 0 0 12 22c4.879 0 8.941-3.494 9.823-8.116c.116-.61.177-1.24.177-1.884a9.98 9.98 0 0 0-3.682-7.752A9.959 9.959 0 0 0 12 2Z"/><path d="M2.085 13.312a6 6 0 0 1 7.297 8.342m12.441-7.769a6.002 6.002 0 0 1-3.505-9.637" opacity=".5"/><path d="M16 16a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-3-7.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/></g>`),
		g.Group(children),
	)
}