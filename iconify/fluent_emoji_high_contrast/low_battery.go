package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LowBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15 9.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5V11h1.5a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5H17v1.5a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5V13h-1.5a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5H15V9.5Z"/><path d="M13.64 1c-.962 0-1.74.778-1.74 1.74v.25H9.82a2.767 2.767 0 0 0-2.77 2.77v22.48a2.767 2.767 0 0 0 2.77 2.77h12.46a2.767 2.767 0 0 0 2.77-2.77V5.76a2.767 2.767 0 0 0-2.77-2.77H20.2v-.25c0-.962-.778-1.74-1.74-1.74h-4.82Zm.26 3.99V3h4.3v1.99h4.08c.428 0 .77.342.77.77V7h-14V5.76c0-.428.342-.77.77-.77h4.08ZM9.05 21V8h14v13h-14Zm0 7.24V27h14v1.24c0 .428-.342.77-.77.77H9.82a.767.767 0 0 1-.77-.77ZM13 23.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-5a.5.5 0 0 1-.5-.5v-1Z"/></g>`),
		g.Group(children),
	)
}