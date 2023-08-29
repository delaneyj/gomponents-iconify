package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drupal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M309 87q-10-6-33-16.5T241 52q-15-10-49-50q-5 36-24 52q-14 11-56 33q-11 6-24.5 16.5t-34 31.5t-34 55T6 264q0 84 60.5 141T210 462t141.5-54.5T410 267q0-115-101-180zm1 331q-15 15-46 18q-56 5-76-14q-6-5 0-9q3-3 6-3q2 0 4 1q15 12 51 12t57-14q6-4 6 1q2 5-2 8zm-77-38q14-12 21-14q5-3 19-3t20 4q5 4 10 16q3 6-4 9q-3 2-6-5q-8-12-22-12q-12 0-28 12q-10 8-13 5q-3-6 3-12zm161-84q0 32-15 54q-15 21-30 20q-7 0-32.5-25.5T280 318q-10 0-55.5 27T154 372q-28 0-43-10q-20-13-20-43q1-27 22.5-47t51.5-20q29-1 67 25t48 26q11 0 43-22t43-22q28 0 28 37z"/>`),
		g.Group(children),
	)
}