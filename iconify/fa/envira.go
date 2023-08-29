package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envira(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M896 816Q792 620 736 538Q597 336 389 220q-34-19-70-36q-89-40-94-32t34 38l39 31q62 43 112.5 93.5T505 431t70.5 113T646 675q9 17 13 25q44 84 84 153t98 154t115.5 150t131 123.5T1236 1371q153 66 154 60q1-3-49-37q-53-36-81-57q-77-58-179-211T896 816zm-347 543q-76-60-132.5-125t-98-143.5t-71-154.5T189 750t-52-209t-60.5-252T0 0q273 0 497.5 36t379 92t271 144.5T1333 445t110 198.5t56 199.5t12.5 198.5t-9.5 173t-20 143.5t-13 107l323 327h-104l-281-285q-22 2-91.5 14t-121.5 19t-138 6t-160.5-17t-167.5-59t-179-111z"/>`),
		g.Group(children),
	)
}