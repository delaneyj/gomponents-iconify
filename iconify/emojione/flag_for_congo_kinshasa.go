package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCongoKinshasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#42ade2" d="M10.8 10.8C2.1 19.5-.1 32.2 4.1 43L43 4.1c-10.8-4.2-23.5-2-32.2 6.7zm42.4 42.4c8.7-8.7 10.9-21.4 6.7-32.2L21 59.9c10.8 4.2 23.5 2 32.2-6.7"/><path fill="#c94747" d="M53.2 10.8c-2-2-4.1-3.6-6.4-4.9L5.9 46.8c1.3 2.3 2.9 4.5 4.9 6.4c2 2 4.1 3.6 6.4 4.9l40.9-40.9c-1.3-2.3-2.9-4.5-4.9-6.4"/><path fill="#ffce31" d="M17.2 58.1c.6.3 1.2.7 1.9 1c.6.3 1.3.6 2 .9l38.8-39c-.3-.7-.5-1.3-.9-2c-.3-.6-.6-1.2-1-1.9l-40.8 41M44.9 4.9c-.7-.3-1.3-.6-2-.9L4.1 43c.3.7.5 1.3.9 2c.3.6.6 1.2 1 1.9l40.8-41c-.6-.4-1.2-.7-1.9-1M18 19.2l3.7 2.8l-1.4-4.6l3.7-2.9h-4.6L18 10l-1.4 4.5H12l3.7 2.9l-1.4 4.6z"/>`),
		g.Group(children),
	)
}