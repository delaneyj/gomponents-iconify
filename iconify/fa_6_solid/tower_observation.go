package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TowerObservation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M241.7 3.4c9-4.5 19.6-4.5 28.6 0l160 80c15.8 7.9 22.2 27.1 14.3 42.9C439 137.5 427.7 144 416 144v80c0 17.7-14.3 32-32 32h-4.9l32 192H480c17.7 0 32 14.3 32 32s-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32h68.9l32-192H128c-17.7 0-32-14.3-32-32v-80c-11.7 0-23-6.5-28.6-17.7c-7.9-15.8-1.5-35 14.3-42.9l160-80zM314.5 448L256 399.2L197.5 448h117zM197.8 256l-4.7 28.3l62.9 52.5l62.9-52.5l-4.7-28.3H197.8zm-13.9 83.2l-11.2 67l45.8-38.2l-34.6-28.8zM293.5 368l45.8 38.1l-11.2-67l-34.6 28.9zM176 128c-8.8 0-16 7.2-16 16s7.2 16 16 16h160c8.8 0 16-7.2 16-16s-7.2-16-16-16H176z"/>`),
		g.Group(children),
	)
}