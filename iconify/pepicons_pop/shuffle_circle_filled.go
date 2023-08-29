package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopShuffleCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6 8.75a1 1 0 0 1 1-1c3.335 0 6.5 2.126 6.5 5.25s-3.165 5.25-6.5 5.25a1 1 0 1 1 0-2c2.74 0 4.5-1.68 4.5-3.25S9.74 9.75 7 9.75a1 1 0 0 1-1-1Z"/><path d="M19 8.75a1 1 0 0 0-1-1c-3.335 0-6.5 2.126-6.5 5.25s3.165 5.25 6.5 5.25a1 1 0 1 0 0-2c-2.74 0-4.5-1.68-4.5-3.25s1.76-3.25 4.5-3.25a1 1 0 0 0 1-1Z"/><path d="M19.46 17.778a.75.75 0 0 1-1.053-.133l-1.5-1.936a.75.75 0 0 1 1.186-.918l1.5 1.935a.75.75 0 0 1-.134 1.052Z"/><path d="M17.05 19.785a.75.75 0 0 1-.15-1.05l1.5-2a.75.75 0 1 1 1.2.9l-1.5 2a.75.75 0 0 1-1.05.15Zm2.41-11.563a.75.75 0 0 0-1.053.133l-1.5 1.936a.75.75 0 0 0 1.186.918l1.5-1.935a.75.75 0 0 0-.134-1.052Z"/><path d="M17.05 6.15a.75.75 0 0 0-.15 1.05l1.5 2a.75.75 0 1 0 1.2-.9l-1.5-2a.75.75 0 0 0-1.05-.15Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopShuffleCircleFilled0)"/></g>`),
		g.Group(children),
	)
}