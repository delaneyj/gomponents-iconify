package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M703.5 384q-26.5 0-45-18.5T640 320q0-34-37.5-64T506 209t-122-17t-122 17t-96.5 47t-37.5 64t37.5 64t96.5 47t122 17q104 0 192.5 34.5t140 93T768 704q0 94-91.5 165.5T448 956v4q0 27-19 45.5t-45.5 18.5t-45-18.5T320 960v-4q-137-15-228.5-86.5T0 704q0-26 18.5-45t45-19t45.5 19t19 45q0 34 37.5 64t96.5 47t122 17t122-17t96.5-47t37.5-64t-37.5-64t-96.5-47t-122-17q-104 0-192.5-34t-140-93T0 320q0-94 91.5-165.5T320 68v-4q0-26 18.5-45t45-19T429 19t19 45v4q137 15 228.5 86.5T768 320q0 27-19 45.5T703.5 384z"/>`),
		g.Group(children),
	)
}