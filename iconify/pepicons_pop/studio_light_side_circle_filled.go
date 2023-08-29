package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioLightSideCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopStudioLightSideCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M16.89 7.456C18.22 6.57 20 7.523 20 9.12v3.263c0 1.597-1.78 2.55-3.11 1.664l-2.945-1.963l.888-1.332l-.888-1.333l2.946-1.963Zm-1.337 3.296L18 12.383V9.12l-2.447 1.632Zm1.361 10.406a1 1 0 0 1-1.32.507L12.5 20.29v.962a1 1 0 1 1-2 0v-.962l-3.094 1.375a1 1 0 1 1-.812-1.827l3.906-1.736v-3.85a1 1 0 0 1 2 0v3.85l3.906 1.736a1 1 0 0 1 .508 1.32Z"/><path d="M7.5 8.752a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4Zm6 0h-4v4h4v-4Z"/><path d="M13.793 8.959a1 1 0 0 1 0-1.415l3-3a1 1 0 1 1 1.414 1.415l-3 3a1 1 0 0 1-1.414 0Zm4.414 8a1 1 0 0 1-1.414 0l-3-3a1 1 0 0 1 1.414-1.415l3 3a1 1 0 0 1 0 1.415Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopStudioLightSideCircleFilled0)"/></g>`),
		g.Group(children),
	)
}