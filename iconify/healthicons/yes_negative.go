package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YesNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsYesNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0Zm-4 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20Zm-9.33-7.741a1 1 0 0 1 .072 1.412l-12.667 14l-.69.761l-.742-.709l-7.334-7a1 1 0 0 1 1.382-1.446l6.59 6.29L33.258 16.33a1 1 0 0 1 1.413-.07Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsYesNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}