package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FieldHockey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#fluentEmojiHighContrastFieldHockey0)"><path d="m12.952 20.698l7.77-18.928C21.329.262 23.006-.343 24.398.195l.02.007l1.975.81l.006.004c1.412.59 2.151 2.227 1.537 3.708l-9.12 22.228a8.219 8.219 0 0 1-10.718 4.479l-.003-.001a8.219 8.219 0 0 1-4.478-10.718l.69-1.68v-.001a4.496 4.496 0 0 1 5.858-2.451a4.496 4.496 0 0 1 2.788 4.118Zm5.311-7.672l-.45 1.097l3.513 1.44l.442-1.076l-3.505-1.46Zm-1.02 2.484l-1.502 3.661l3.514 1.441l1.502-3.661l-3.513-1.44Zm-2.45 5.97l-1.117 2.72c-.38.93-1.45 1.38-2.38.99a1.817 1.817 0 0 1-.99-2.38l.46-1.13c.52-1.27-.09-2.73-1.36-3.25s-2.73.09-3.25 1.36l-.69 1.68c-1.3 3.17.21 6.8 3.39 8.11c3.17 1.3 6.8-.21 8.11-3.39l1.341-3.268l-3.514-1.442ZM27 27a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g><defs><clipPath id="fluentEmojiHighContrastFieldHockey0"><path fill="#fff" d="M0 0h32v32H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}