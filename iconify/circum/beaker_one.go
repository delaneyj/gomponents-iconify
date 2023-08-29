package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeakerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.447 18.645l-.51-1.52a17.9 17.9 0 0 0-4.02-6.66a1.493 1.493 0 0 1-.42-1.04v-6.36H15a.5.5 0 0 0 0-1H9a.5.5 0 0 0 0 1h.5v6.36a1.484 1.484 0 0 1-.41 1.04a17.9 17.9 0 0 0-4.02 6.66l-.52 1.52a2.5 2.5 0 0 0 2.37 3.29h10.16a2.5 2.5 0 0 0 2.37-3.29Zm-9.64-7.49a2.477 2.477 0 0 0 .69-1.73v-6.36h3v6.36a2.486 2.486 0 0 0 .7 1.73a16.907 16.907 0 0 1 3.01 4.38H6.787a16.943 16.943 0 0 1 3.02-4.38Zm8.49 9.16a1.507 1.507 0 0 1-1.22.62H6.917a1.5 1.5 0 0 1-1.42-1.98l.51-1.52q.15-.45.33-.9h11.32q.18.45.33.9l.51 1.52a1.5 1.5 0 0 1-.197 1.36Z"/>`),
		g.Group(children),
	)
}