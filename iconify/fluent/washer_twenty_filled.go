package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WasherTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.484 10.668a3.59 3.59 0 0 0-.483.262a3 3 0 0 1 5.846-.878c-.393.399-1.263.829-2.726.463c-1.104-.276-1.969-.144-2.637.153ZM6 3a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Zm4 4a4 4 0 0 1 3.922 3.211l.025.013l-.016.031A4 4 0 1 1 10 7ZM7.5 5.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm3-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}