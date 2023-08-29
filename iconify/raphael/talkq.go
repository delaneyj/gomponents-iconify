package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Talkq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4.938c-7.732 0-14 4.7-14 10.5c0 1.98.74 3.833 2.016 5.414L2 25.272l5.613-1.44c2.34 1.316 5.237 2.106 8.387 2.106c7.732 0 14-4.7 14-10.5s-6.268-10.5-14-10.5zm.868 16.437h-1.97v-1.89h1.97v1.89zm-.096-3.28h-1.777l-.176-8.084h2.112l-.16 8.084z"/>`),
		g.Group(children),
	)
}