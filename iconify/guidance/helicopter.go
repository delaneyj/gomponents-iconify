package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Helicopter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15.5 7.5H5.158a5.857 5.857 0 0 1-4.408-2H.5v5h4a7 7 0 0 1 7 7h12v-2a8 8 0 0 0-8-8Zm0 0v3a4 4 0 0 0 4 4h3.938m-7.938-7V4m0 0c-2.275 0-4.544-.245-6.767-.73L7.5 3m8 1c2.275 0 4.544-.245 6.767-.73L23.5 3m-9 14.5v3.754m6-3.754v3.754m-9 .746l1.096-.313a17.851 17.851 0 0 1 9.808 0L23.5 22"/>`),
		g.Group(children),
	)
}