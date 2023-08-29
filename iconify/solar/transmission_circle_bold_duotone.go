package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransmissionCircleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z" opacity=".5"/><path d="M8 8.25a.75.75 0 0 1 .75.75v2.25h2.5V9a.75.75 0 0 1 1.5 0v2.25H13c.476 0 .796 0 1.043-.017c.241-.017.358-.046.435-.078c.307-.127.55-.37.677-.677c.032-.077.061-.194.078-.435c.017-.247.017-.567.017-1.043a.75.75 0 0 1 1.5 0v.025c0 .445 0 .816-.02 1.12a2.822 2.822 0 0 1-.19.907a2.75 2.75 0 0 1-1.488 1.489c-.29.12-.59.167-.907.188c-.304.021-.675.021-1.12.021h-.275V15a.75.75 0 0 1-1.5 0v-2.25h-2.5V15a.75.75 0 0 1-1.5 0V9A.75.75 0 0 1 8 8.25Z"/></g>`),
		g.Group(children),
	)
}