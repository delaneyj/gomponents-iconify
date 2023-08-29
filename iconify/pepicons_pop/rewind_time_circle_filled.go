package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindTimeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopRewindTimeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.955 8.805a5.94 5.94 0 0 0-1.736 4.87a1 1 0 0 1-1.99.21a7.94 7.94 0 0 1 2.322-6.504c3.136-3.092 8.19-3.04 11.289.102c3.1 3.143 3.081 8.198-.054 11.29a7.924 7.924 0 0 1-5.44 2.286a8.021 8.021 0 0 1-2.283-.29a1 1 0 1 1 .533-1.927a6 6 0 0 0 1.714.217a5.924 5.924 0 0 0 4.071-1.71c2.343-2.31 2.365-6.099.035-8.461c-2.33-2.363-6.118-2.393-8.46-.083Z"/><path d="M6.967 14.695a1 1 0 0 1-1.412.082l-1.72-1.53a1 1 0 0 1 1.33-1.494l1.72 1.53a1 1 0 0 1 .082 1.412Z"/><path d="M5.42 14.6a1 1 0 0 0 1.4.2l2-1.5a1 1 0 0 0-1.2-1.6l-2 1.5a1 1 0 0 0-.2 1.4ZM13 9a1 1 0 0 1 1 1v3.5a1 1 0 1 1-2 0V10a1 1 0 0 1 1-1Z"/><path d="M17 13.5a1 1 0 0 1-1 1h-3a1 1 0 1 1 0-2h3a1 1 0 0 1 1 1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopRewindTimeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}