package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButtonOnFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaRadioButtonOnFill0"><g id="evaRadioButtonOnFill1"><g id="evaRadioButtonOnFill2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M12 7a5 5 0 1 0 5 5a5 5 0 0 0-5-5Z"/></g></g></g>`),
		g.Group(children),
	)
}