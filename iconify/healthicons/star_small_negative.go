package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarSmallNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsStarSmallNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24.917 14.568a1.025 1.025 0 0 0-1.834 0L20.63 19.52c-.15.3-.438.51-.77.558l-5.484.793a1.018 1.018 0 0 0-.567 1.739l3.968 3.854c.24.234.35.571.294.902l-.937 5.441c-.143.833.734 1.468 1.484 1.075l4.905-2.57c.298-.156.654-.156.952 0l4.905 2.57c.75.393 1.627-.242 1.484-1.075l-.937-5.441c-.057-.33.053-.668.294-.902l3.968-3.854a1.018 1.018 0 0 0-.567-1.739l-5.483-.794a1.023 1.023 0 0 1-.77-.557l-2.453-4.95Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsStarSmallNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}