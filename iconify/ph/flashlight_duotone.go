package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 64v13.33a8 8 0 0 1-1.6 4.8l-20.8 27.74a8 8 0 0 0-1.6 4.8V224a8 8 0 0 1-8 8H96a8 8 0 0 1-8-8V114.67a8 8 0 0 0-1.6-4.8L65.6 82.13a8 8 0 0 1-1.6-4.8V64Z" opacity=".2"/><path d="M184 16H72a16 16 0 0 0-16 16v45.33a16.12 16.12 0 0 0 3.2 9.6L80 114.67V224a16 16 0 0 0 16 16h64a16 16 0 0 0 16-16V114.67l20.8-27.74a16.12 16.12 0 0 0 3.2-9.6V32a16 16 0 0 0-16-16ZM72 32h112v24H72V32Zm91.2 73.07a16.12 16.12 0 0 0-3.2 9.6V224H96V114.67a16.12 16.12 0 0 0-3.2-9.6L72 77.33V72h112v5.33ZM136 120v32a8 8 0 0 1-16 0v-32a8 8 0 0 1 16 0Z"/></g>`),
		g.Group(children),
	)
}