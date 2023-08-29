package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipNoteO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1664 256v637q0 110-.5 153t-1.5 130t-3 122.5t-6.5 102t-11 97.5t-16 80.5t-22 80t-29.5 66t-38 67.5H0q79-242 103.5-447.5T128 768V256h256v-54q0-88 68-147Q517 0 617 0q91 0 151.5 39t80.5 88q13 33 13 89H734q-5-58-35-82.5T609 109q-52 0-82 50q-6 11-9.5 25t-4.5 21t-1 27v24h1152zM512 384v448q-1 32 15 46.5t49 17.5q40 0 52-12.5t12-51.5V512h128v320q0 71-12 99q-41 93-180 93q-86 0-139-53t-53-139V384H256v384q0 326-17.5 515T167 1664h1275q28-57 46-122.5t27.5-118.5t14-155t5.5-163.5t1-211.5V384H512z"/>`),
		g.Group(children),
	)
}