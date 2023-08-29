package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeKnobOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11.25 7.816a4.251 4.251 0 1 0 1.5 0V11a.75.75 0 0 1-1.5 0V7.816ZM6.25 12a5.75 5.75 0 1 1 11.5 0a5.75 5.75 0 0 1-11.5 0Z" clip-rule="evenodd"/><path d="M13 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm7.5 9.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-17 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm3.197-7.718a1 1 0 1 1-1.415 1.415a1 1 0 0 1 1.415-1.415Zm12.021 12.021a1 1 0 1 1-1.415 1.415a1 1 0 0 1 1.415-1.415Zm0-10.606a1 1 0 1 1-1.415-1.415a1 1 0 0 1 1.415 1.415ZM6.697 18.718a1 1 0 1 1-1.415-1.415a1 1 0 0 1 1.415 1.415Z"/></g>`),
		g.Group(children),
	)
}