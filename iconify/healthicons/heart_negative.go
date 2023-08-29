package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsHeartNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM15.562 7C10.037 7 6 12.64 6 18.724C6 32.304 24 41 24 41s18-9.256 18-22.276C42 12.642 37.965 7 32.437 7c-3.835 0-6.68 2.531-8.437 6.121C22.243 9.531 19.398 7 15.562 7Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHeartNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}