package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunningShirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#6adbc6" d="M49.4 4.7C49.2 3.1 49 1.5 49 0h-3.9c.1.5.1 1 .1 1.6c0 6.5-5.9 11.8-13.2 11.8c-7.3 0-13.2-5.3-13.2-11.8c0-.5 0-1.1.1-1.6H15c0 1.5-.2 3.1-.4 4.7c-1.7 9.7-4.1 17.6-10.7 19.5c0 0 2.8 4.5 2.8 12.3C6.7 50 3.9 64 3.9 64h56.2s-3.5-14-3.5-27.9c0-7.4 3.5-11.9 3.5-11.9c-6.5-1.9-9-9.8-10.7-19.5"/><path fill="#428bc1" d="M32 13.4c7.3 0 13.2-5.3 13.2-11.8c0-.5 0-1.1-.1-1.6c-3.6 1.6-7 3.7-13.1 3.7c-6.2 0-9.4-2-13.1-3.7c-.1.5-.1 1-.1 1.6c0 6.5 5.9 11.8 13.2 11.8"/><path fill="#ffce31" d="M5.4 54.8c-.8 6-1.5 9.2-1.5 9.2h8.3l42.9-43.1c-2.2-2.6-3.5-6.3-4.6-10.6L5.4 54.8"/>`),
		g.Group(children),
	)
}