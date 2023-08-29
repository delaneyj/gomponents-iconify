package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M678 832H346q-47-33-89-76q-34 12-65 12q-80 0-136-56.5T0 576q0-63 37-112.5t95-69.5q1-6 1-10h758q5 35 5 128q0 45-21 94t-55 91t-69.5 75.5T678 832zM128 512q0-15 1-47q-30 17-47.5 46.5T64 576q0 53 37.5 90.5T192 704q4 0 16-2q-80-101-80-190zm448-320q0 23 9 55q39 28 51 73h-64q-2-14-5-24q-55-38-55-104q0-53 37.5-90.5T640 64q15 0 32 4q-42 11-69 45t-27 79zm-192-64q0 23 9 55q55 39 55 105q0 14-5 32h-63q4-18 4-32q0-24-9-56q-55-38-55-104q0-53 37.5-90.5T448 0q15 0 32 4q-42 11-69 45t-27 79zm384 896H256q-55 0-116-21.5t-100.5-52T0 896h1024q0 24-39.5 54.5t-100.5 52t-116 21.5z"/>`),
		g.Group(children),
	)
}