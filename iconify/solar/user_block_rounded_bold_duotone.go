package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserBlockRoundedBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="6" r="4"/><path fill-rule="evenodd" d="M18 15.75a2.25 2.25 0 0 0-2.03 3.22l3-3a2.24 2.24 0 0 0-.97-.22Zm2.03 1.28l-3 3a2.25 2.25 0 0 0 3-3Zm-5.78.97a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Z" clip-rule="evenodd"/><path d="M17.216 14.332a3.751 3.751 0 0 0-1.97 6.213c-.97.29-2.075.455-3.246.455c-3.866 0-7-1.79-7-4s3.134-4 7-4c2.072 0 3.934.514 5.216 1.332Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}