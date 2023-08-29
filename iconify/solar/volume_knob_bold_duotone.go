package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeKnobBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.25 7.056a5.001 5.001 0 1 0 1.5 0V11a.75.75 0 0 1-1.5 0V7.056ZM13 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M20.5 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-17 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z" opacity=".5"/><path d="M6.696 5.282a1 1 0 1 1-1.414 1.415a1 1 0 0 1 1.414-1.415Z" opacity=".7"/><path d="M18.718 17.303a1 1 0 1 1-1.414 1.415a1 1 0 0 1 1.414-1.415Z" opacity=".4"/><path d="M18.718 6.697a1 1 0 1 1-1.414-1.415a1 1 0 0 1 1.414 1.415Z" opacity=".7"/><path d="M6.696 18.718a1 1 0 1 1-1.414-1.415a1 1 0 0 1 1.414 1.415Z" opacity=".4"/></g>`),
		g.Group(children),
	)
}