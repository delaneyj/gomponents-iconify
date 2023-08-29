package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeigherLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 10c0-3.771 0-5.657 1.172-6.828C5.343 2 7.229 2 11 2h2c3.771 0 5.657 0 6.828 1.172C21 4.343 21 6.229 21 10v4c0 3.771 0 5.657-1.172 6.828C18.657 22 16.771 22 13 22h-2c-3.771 0-5.657 0-6.828-1.172C3 19.657 3 17.771 3 14v-4Z"/><path stroke-linecap="round" d="M8 18h8m.412-8.03l.665-1.663A2 2 0 0 0 15.55 5.59l-.425-.07a19 19 0 0 0-6.247 0l-.426.07a2 2 0 0 0-1.528 2.716L7.59 9.97c.235.59.87.914 1.486.76c1.92-.48 3.93-.48 5.85 0a1.27 1.27 0 0 0 1.487-.76Zm-6.232-.041l-.676-1.883"/></g>`),
		g.Group(children),
	)
}