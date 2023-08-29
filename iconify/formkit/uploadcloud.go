package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uploadcloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 11c-.28 0-.5-.22-.5-.5s.22-.5.5-.5c1.65 0 3-1.35 3-3s-1.35-3-3-3h-1.05c-.18 0-.34-.09-.43-.25C9.88 2.65 8.76 2 7.51 2c-1.93 0-3.5 1.57-3.5 3.5c0 .28-.22.5-.5.5h-.5c-1.1 0-2 .9-2 2s.9 2 2 2c.28 0 .5.22.5.5s-.22.5-.5.5c-1.65 0-3-1.35-3-3s1.35-3 3-3h.03c.25-2.25 2.16-4 4.47-4c1.49 0 2.89.76 3.72 2h.78c2.21 0 4 1.79 4 4s-1.79 4-4 4Z"/><path fill="currentColor" d="M10 9.25a.47.47 0 0 1-.35-.15L7.5 6.95L5.35 9.1c-.2.2-.51.2-.71 0s-.2-.51 0-.71l2.5-2.5c.2-.2.51-.2.71 0l2.5 2.5c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M7.5 13c-.28 0-.5-.22-.5-.5v-6c0-.28.22-.5.5-.5s.5.22.5.5v6c0 .28-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}