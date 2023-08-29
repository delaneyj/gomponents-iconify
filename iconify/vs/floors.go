package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Floors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M33 431q-18-9-25.5-19.5T0 383t7.5-28.5T33 335L670 17q39-19 98-17q59-2 98 17l637 318q18 9 25.5 19.5t7.5 28.5t-7.5 28.5T1503 431L866 749q-39 19-98 17q-59 2-98-17zm0 770q-18-9-25.5-19.5T0 1153t7.5-28.5T33 1105l160-80l477 238q40 19 98 16q58 3 98-16l477-238l160 80q18 9 25.5 19.5t7.5 28.5t-7.5 28.5t-25.5 19.5l-637 318q-40 19-98 16q-58 3-98-16zm0-384q-18-9-25.5-19.5T0 769t7.5-28.5T33 721l160-80l477 238q40 19 98 16q58 3 98-16l477-238l160 80q18 9 25.5 19.5t7.5 28.5t-7.5 28.5T1503 817l-637 318q-40 19-98 16q-58 3-98-16z"/>`),
		g.Group(children),
	)
}