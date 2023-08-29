package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundcloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M9 260H0v-51h9v51zm21 17h-9v-81h9v81zm17 4h-9v-94h9v94zm17 4h-9v-94h9v94zm21 0h-8V162h8v123zm17 0h-8V145h8v140zm22 0h-9V136h9v149zm17 0h-9V132h9v153zm21 0h-8V136h8v149zm17 0h-8V140h8v145zm17 0h-8V123h8v162zm22 0h-9V110h9v175zm156-1H228q-6 0-6-6V111q0-4 5-6q17-6 34-6q36 0 62.5 24.5T354 183q9-4 20-4q22 0 37.5 15.5T427 232t-15.5 37t-37.5 15z"/>`),
		g.Group(children),
	)
}