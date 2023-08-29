package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevenMind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.528 16.431V31.57a4.362 4.362 0 0 0 2.181 3.778l13.11 7.569a4.362 4.362 0 0 0 4.362 0l13.11-7.57a4.362 4.362 0 0 0 2.181-3.777V16.43a4.362 4.362 0 0 0-2.18-3.778l-13.11-7.569a4.362 4.362 0 0 0-4.363 0l-13.11 7.57a4.362 4.362 0 0 0-2.18 3.777Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.304 41.408l9.334-17.347H17.146"/>`),
		g.Group(children),
	)
}