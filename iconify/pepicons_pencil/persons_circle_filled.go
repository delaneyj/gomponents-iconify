package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonsCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPersonsCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.36 12.977a5.5 5.5 0 0 0-.923 3.05V17a.5.5 0 1 1-1 0v-.972a6.5 6.5 0 0 1 1.092-3.606l.108-.162a.5.5 0 1 1 .832.555l-.108.162Z"/><path d="M9.18 11.365c-1.09 0-2.107.544-2.711 1.45l-.832-.554a4.258 4.258 0 0 1 3.542-1.896h.22a.5.5 0 0 1 0 1h-.22Zm3.078 1.6c.47.706.721 1.534.721 2.382h1a5.289 5.289 0 0 0-.889-2.936l-.1-.15a.5.5 0 1 0-.832.554l.1.15Z"/><path d="M9.448 11.365c1.089 0 2.106.544 2.71 1.45l.832-.554a4.258 4.258 0 0 0-3.542-1.896h-.22a.5.5 0 1 0 0 1h.22Z"/><path d="M9.25 10.25a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5Zm0 1a3.25 3.25 0 1 0 0-6.5a3.25 3.25 0 0 0 0 6.5Zm4.258 4.936a5.5 5.5 0 0 0-.923 3.051v1.034a.5.5 0 1 1-1 0v-1.034a6.5 6.5 0 0 1 1.091-3.605l.133-.2a.5.5 0 0 1 .832.556l-.133.198Z"/><path d="M16.42 14.5a3.34 3.34 0 0 0-2.78 1.488l-.831-.555A4.34 4.34 0 0 1 16.42 13.5h.224a.5.5 0 1 1 0 1h-.224Zm3.187 1.686a5.5 5.5 0 0 1 .924 3.051v1.034a.5.5 0 1 0 1 0v-1.034a6.5 6.5 0 0 0-1.092-3.605l-.133-.2a.5.5 0 1 0-.832.556l.133.198Z"/><path d="M16.695 14.5a3.34 3.34 0 0 1 2.78 1.488l.832-.555a4.34 4.34 0 0 0-3.612-1.933h-.225a.5.5 0 1 0 0 1h.225Z"/><path d="M16.5 13.5a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5Zm0 1a3.25 3.25 0 1 0 0-6.5a3.25 3.25 0 0 0 0 6.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPersonsCircleFilled0)"/></g>`),
		g.Group(children),
	)
}