package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Talke(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4.938c-7.732 0-14 4.7-14 10.5c0 1.98.74 3.833 2.016 5.414L2 25.272l5.613-1.44c2.34 1.316 5.237 2.106 8.387 2.106c7.732 0 14-4.7 14-10.5s-6.268-10.5-14-10.5zm.982 16.437h-1.97v-1.89h1.97v1.89zm0-3.906v.624h-1.97v-.77c0-2.32 2.642-2.688 2.642-4.336c0-.752-.672-1.33-1.553-1.33c-.91 0-1.712.673-1.712.673l-1.12-1.392s1.104-1.153 3.01-1.153c1.81 0 3.49 1.12 3.49 3.01c0 2.642-2.786 2.946-2.786 4.674z"/>`),
		g.Group(children),
	)
}