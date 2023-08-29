package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Auto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#FAB431"/><path fill="#FFF" d="m24.5 19.267l-4.185-2.511l.48-1.989l1.787 1.055V12.19l-5.74-3.422v7.9l-1.365.755l7.538 4.47L16 26l-8.5-4.978v-9.955L16 6l8.5 5.067v8.2zM9.418 12.19v6.478l5.518-3.2V8.9l-5.518 3.289zm4.097 6.344l-3.138 1.919L16 23.745l3.175-1.86l-5.66-3.352z"/></g>`),
		g.Group(children),
	)
}