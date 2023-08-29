package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBotswana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#b4d7ee" d="M60.9 24C57.4 11.3 45.8 2 32 2S6.6 11.3 3.1 24v16C6.6 52.7 18.2 62 32 62c13.8 0 25.4-9.3 28.9-22V24"/><path fill="#fff" d="M61.6 27c-.2-1-.4-2-.7-3H3.1c-.3 1-.5 2-.7 3h59.2M2.4 37c.2 1 .4 2 .7 3h57.8c.3-1 .5-2 .7-3H2.4"/><path fill="#3e4347" d="M2 32c0 1.7.1 3.4.4 5h59.2c.3-1.6.4-3.3.4-5s-.1-3.4-.4-5H2.4c-.3 1.6-.4 3.3-.4 5"/>`),
		g.Group(children),
	)
}