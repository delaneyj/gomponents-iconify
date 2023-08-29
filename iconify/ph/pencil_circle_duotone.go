package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M176 211.16Zm-96 0ZM224 128a96 96 0 1 0-144 83.16V176l48-104l48 104v35.16A96 96 0 0 0 224 128Z" opacity=".2"/><path d="M201.54 54.46A104 104 0 0 0 54.46 201.54A104 104 0 0 0 201.54 54.46ZM88 192a16 16 0 0 1 32 0v23.59a88 88 0 0 1-32-9.22Zm48 0a16 16 0 0 1 32 0v14.37a88 88 0 0 1-32 9.22Zm-28.73-56h41.46l11.58 25.1a31.93 31.93 0 0 0-32.31 9.77a31.93 31.93 0 0 0-32.31-9.77Zm7.39-16L128 91.09L141.34 120Zm75.56 70.23c-2 2-4.08 3.87-6.22 5.64V176a7.91 7.91 0 0 0-.74-3.35l-48-104a8 8 0 0 0-14.52 0l-48 104A7.91 7.91 0 0 0 72 176v19.87a88.917 88.917 0 0 1-6.22-5.64a88 88 0 1 1 124.44 0Z"/></g>`),
		g.Group(children),
	)
}