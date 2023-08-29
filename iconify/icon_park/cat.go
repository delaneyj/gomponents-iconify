package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M42 26C42 34.8366 33.9411 42 24 42C14.0589 42 6 34.8366 6 26M15 12.1405C17.6476 10.7792 20.7214 10 24 10C27.2786 10 30.3524 10.7792 33 12.1405"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 26V8.48814C6 6.757 8.05005 5.84346 9.33729 7.00098L15 12.093"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 26V8.48814C42 6.757 39.9499 5.84346 38.6627 7.00098L33 12.093"/><circle cx="30" cy="22" r="2" fill="#000"/><circle cx="18" cy="22" r="2" fill="#000"/><circle cx="24" cy="28" r="2" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 30L4 31"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 35L7 41"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 30L44 31"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M29 35L41 41"/></g>`),
		g.Group(children),
	)
}