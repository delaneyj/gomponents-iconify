package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 128a96 96 0 1 1-96-96a96 96 0 0 1 96 96Z" opacity=".2"/><path d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm87.63 96H175.8c-1.41-28.46-10.27-55.47-25.12-77a88.2 88.2 0 0 1 64.95 77Zm-119.4 16h63.54c-1.68 30.87-13 59.62-31.77 79.89c-18.73-20.27-30.09-49.02-31.77-79.89Zm0-16c1.68-30.87 13-59.62 31.77-79.89c18.73 20.27 30.09 49 31.77 79.89Zm9.09-77C90.47 64.53 81.61 91.54 80.2 120H40.37a88.2 88.2 0 0 1 64.95-77Zm-64.95 93H80.2c1.41 28.46 10.27 55.47 25.12 77a88.2 88.2 0 0 1-64.95-77Zm110.31 77c14.85-21.56 23.71-48.57 25.12-77h39.83a88.2 88.2 0 0 1-64.95 77Z"/></g>`),
		g.Group(children),
	)
}