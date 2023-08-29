package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13.464 20.535a5 5 0 1 0 7.072-7.07a5 5 0 0 0-7.072 7.07ZM3 7a5 5 0 1 0 10 0A5 5 0 0 0 3 7Z" opacity=".5"/><path d="M12.056 17.75h9.888a5.026 5.026 0 0 0 0-1.5h-9.888a5.026 5.026 0 0 0 0 1.5ZM3.974 4.035l6.991 6.991a5.038 5.038 0 0 0 1.06-1.06l-6.99-6.992a5.028 5.028 0 0 0-1.061 1.06Z"/></g>`),
		g.Group(children),
	)
}