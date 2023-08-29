package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceShuttle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M620 864q-110 64-268 64H224v-64h-64q-13 0-22.5-23.5T128 784q0-24 7-49q-58-2-96.5-10.5T0 704t38.5-20.5T135 673q-7-25-7-49q0-33 9.5-56.5T160 544h64v-64h128q158 0 268 64h1113q42 7 106.5 18t80.5 14q89 15 150 40.5t83.5 47.5t22.5 40t-22.5 40t-83.5 47.5t-150 40.5q-16 3-80.5 14T1733 864H620zm1119-252q53 36 53 92t-53 92l81 30q68-48 68-122t-68-122zM625 880h1015q-217 38-456 80q-57 0-113 24t-83 48l-28 24l-288 288q-26 26-70.5 45t-89.5 19h-96l-93-464h29q157 0 273-64zM352 464h-29L416 0h96q46 0 90 19t70 45l288 288q4 4 11 10.5t30.5 23t48.5 29t61.5 23T1184 448l456 80H625q-116-64-273-64z"/>`),
		g.Group(children),
	)
}