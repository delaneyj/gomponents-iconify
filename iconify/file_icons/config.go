package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Config(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M36.571 475.429h73.143V365.714H36.571V475.43zM109.714 36.57H36.571v182.86h73.143V36.57zm182.857 0H219.43v73.143h73.142V36.571zM0 329.143h146.286V256H0v73.143zm219.429 146.286h73.142V256H219.43v219.429zm-36.572-256h146.286v-73.143H182.857v73.143zM475.43 36.57h-73.143V256h73.143V36.571zm-109.715 256v73.143H512V292.57H365.714zm36.572 182.858h73.143v-73.143h-73.143v73.143z"/>`),
		g.Group(children),
	)
}