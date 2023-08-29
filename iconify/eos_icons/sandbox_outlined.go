package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SandboxOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="9.537" cy="8.987" r="1" fill="currentColor"/><path fill="currentColor" d="M8.037 10.987a1 1 0 0 0-1 1a.984.984 0 0 0 .482.834a.972.972 0 0 0 .518.166a1 1 0 0 0 1-1a1.016 1.016 0 0 0-1-1Z"/><path fill="currentColor" d="M3 19a2 2 0 0 0 2 2h14a2.006 2.006 0 0 0 2-2V7.244l-2 2V15h-4a3 3 0 0 1-6 0H5L4.99 5h9.769l2-2H5a1.993 1.993 0 0 0-2 2Zm2 0v-2h2.422a5.103 5.103 0 0 0 1.565 2Zm14 0l-3.987-.033A4.998 4.998 0 0 0 16.578 17H19Z"/><path fill="currentColor" d="M13.943 7.071c-.376-.376-.88-.17-1.157.31l-1.164 1.914a2.08 2.08 0 0 1 1.487 1.15a1.618 1.618 0 0 1 .592-.128l.01-.001a1.703 1.703 0 0 1 1.683 1.67l.018.014l.38-.225l.871-.516c.476-.29.687-.78.31-1.157l-.84-.84l5.274-5.274l-1.415-1.415l-5.273 5.274Z"/><circle cx="11.037" cy="11.987" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}