package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feComments0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feComments1" fill="currentColor"><path id="feComments2" d="M4.368 14.632L3 16v-2.8A5.64 5.64 0 0 1 2 10c0-3.314 2.91-6 6.5-6c3.254 0 5.95 2.207 6.425 5.088A6.57 6.57 0 0 1 16 9c3.314 0 6 2.462 6 5.5c0 1.125-.368 2.17-1 3.041V20l-1.225-1.225A6.32 6.32 0 0 1 16 20c-2.825 0-5.194-1.79-5.831-4.2c-.533.13-1.092.2-1.669.2a6.81 6.81 0 0 1-4.132-1.368ZM8.5 14c2.52 0 4.5-1.828 4.5-4c0-2.172-1.98-4-4.5-4S4 7.828 4 10c0 2.172 1.98 4 4.5 4Zm3.546 1.03C12.336 16.687 13.972 18 16 18c2.24 0 4-1.6 4-3.5S18.24 11 16 11c-.389 0-.763.048-1.117.138c-.338 1.626-1.387 3.018-2.837 3.891Z"/></g></g>`),
		g.Group(children),
	)
}