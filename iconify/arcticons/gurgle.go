package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gurgle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.574 32.172v7h3.5m13.365 0h3.501m-3.501-7h3.501m-3.501 3.5h2.282m-2.282-3.5v7M18.341 11.175a2.344 2.344 0 0 0-2.467-2.344a2.434 2.434 0 0 0-2.215 2.476v2.175A2.344 2.344 0 0 0 16 15.828h0a2.344 2.344 0 0 0 2.341-2.346H16m18.341 9.364a2.344 2.344 0 0 0-2.467-2.343a2.434 2.434 0 0 0-2.215 2.476v2.174A2.344 2.344 0 0 0 32 27.5h0a2.344 2.344 0 0 0 2.341-2.346H32M13.655 27.5v-7h2.292a2.35 2.35 0 0 1 0 4.702h-2.292m2.292 0l2.291 2.296m11.417-18.67v4.681a2.319 2.319 0 0 0 4.638 0V8.828"/>`),
		g.Group(children),
	)
}