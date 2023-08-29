package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncomingEnvelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.27 14a2.5 2.5 0 0 0-2.465 2.089l-2 12A2.5 2.5 0 0 0 8.27 31h17.958a2.5 2.5 0 0 0 2.466-2.089l2-12A2.5 2.5 0 0 0 28.23 14H10.271Zm-.493 2.418a.5.5 0 0 1 .494-.418h17.958a.5.5 0 0 1 .494.582l-.094.559l-10.29 5.54a.5.5 0 0 1-.523-.03l-8.129-5.69l.09-.543Zm-.454 2.729l5.16 3.611l-6.33 3.409l1.17-7.02Zm6.994 4.895l.352.247a2.5 2.5 0 0 0 2.62.153l.371-.2l7.023 4.469a.5.5 0 0 1-.454.289H8.271a.5.5 0 0 1-.482-.366l8.528-4.592Zm5.361-.887l6.535-3.519l-1.156 6.942l-5.379-3.423ZM1.75 16a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5h-4.5Zm0 4a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5ZM1 24.75a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}