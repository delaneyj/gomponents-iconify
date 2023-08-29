package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButtonOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaRadioButtonOffFill0"><g id="evaRadioButtonOffFill1"><path id="evaRadioButtonOffFill2" fill="currentColor" d="M12 22a10 10 0 1 1 10-10a10 10 0 0 1-10 10Zm0-18a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/></g></g>`),
		g.Group(children),
	)
}