package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q139 0 90 49Q72 67 68 87q-29 5-48.5 28.5T0 169v174q0 31 18.5 54T66 425q4 19 21 38q46 49 160 49h9q96-1 161-53.5T501 337q11-41 11-81t-11-81q-19-70-85-122.5T256 0zM83 384q-16 0-28-12.5T43 343V169q0-16 12-28.5T83 128h4q17 0 29 12t12 29v174q0 17-12 29t-29 12h-4zm378-58q-8 29-29.5 59T363 441.5T256 469q-105 2-137-34q-5-5-8-13q26-8 43-30t17-49V169q0-28-16-49.5T113 90q2-6 9-11q32-36 134-36q86 0 137.5 46t67.5 97q8 28 8 70t-8 70zm-333-91q0 17-12.5 29.5T85 277q-17 0-29.5-12.5T43 235q0-18 12.5-30.5T85 192q18 0 30.5 12.5T128 235z"/>`),
		g.Group(children),
	)
}