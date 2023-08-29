package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkFiveG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M35.087 9.408a12 12 0 0 0-2.602-3.893L33.9 4.1a14 14 0 0 1 0 19.799l-1.415-1.415a12 12 0 0 0 2.602-13.077Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M29.543 11.704a6 6 0 0 0-1.3-1.947l1.414-1.414a8 8 0 0 1 0 11.314l-1.414-1.414a6 6 0 0 0 1.3-6.539Zm-9.786-1.947a6 6 0 0 0 0 8.486l-1.414 1.414a8 8 0 0 1 0-11.314l1.414 1.414Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.515 5.515a12 12 0 0 0 0 16.97L14.1 23.9a14 14 0 0 1 0-19.8l1.415 1.415Z" clip-rule="evenodd"/><path d="M26 14a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path fill-rule="evenodd" d="M10 29a3 3 0 0 1 3-3h22a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H13a3 3 0 0 1-3-3V29Zm12 1a1 1 0 0 0-1-1h-5.25a1 1 0 0 0-.98.804l-.75 3.75a1 1 0 0 0 1.535 1.028c.21-.14 1.143-.582 2.445-.582a2 2 0 1 1 0 4h-.832c-.577 0-1.054-.36-1.227-.84a1 1 0 1 0-1.882.68A3.297 3.297 0 0 0 17.17 40H18a4 4 0 0 0 0-8c-.61 0-1.173.072-1.666.182L16.57 31H21a1 1 0 0 0 1-1Zm10.753 1.774a1 1 0 0 0-.08-1.412a5.503 5.503 0 1 0 0 8.204a1 1 0 0 0 .334-.745v-3.357a1 1 0 0 0-1-1h-3.002a1 1 0 1 0 0 2h2.002v1.875a3.505 3.505 0 0 1-5.506-2.875a3.503 3.503 0 0 1 5.84-2.611a1 1 0 0 0 1.412-.079Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}