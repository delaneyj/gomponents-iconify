package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldChevronLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 42H48a14 14 0 0 0-14 14v58.77c0 88.24 74.68 117.52 89.65 122.49a13.5 13.5 0 0 0 8.7 0c15-5 89.65-34.25 89.65-122.49V56a14 14 0 0 0-14-14Zm-79.44 183.89a1.55 1.55 0 0 1-1.12 0c-9.46-3.14-45.14-17-66-52L128 127.32l66.56 46.59c-20.86 35.02-56.56 48.84-66 51.98ZM210 114.79c0 19-3.83 35-9.85 48.39l-68.71-48.1a6 6 0 0 0-6.88 0l-68.71 48.1c-6-13.4-9.85-29.38-9.85-48.39V56a2 2 0 0 1 2-2h160a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}