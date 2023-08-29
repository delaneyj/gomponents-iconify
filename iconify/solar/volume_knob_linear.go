package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeKnobLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="5" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 7v4"/><path fill="currentColor" d="M13 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm7.5 9.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-17 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm3.197-7.718a1 1 0 1 1-1.414 1.415a1 1 0 0 1 1.414-1.415Zm12.02 12.021a1 1 0 1 1-1.414 1.415a1 1 0 0 1 1.414-1.415Zm0-10.606a1 1 0 1 1-1.414-1.415a1 1 0 0 1 1.414 1.415ZM6.697 18.718a1 1 0 1 1-1.414-1.415a1 1 0 0 1 1.414 1.415Z"/></g>`),
		g.Group(children),
	)
}