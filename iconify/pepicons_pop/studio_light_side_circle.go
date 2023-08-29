package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioLightSideCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M16.89 7.456C18.22 6.57 20 7.523 20 9.12v3.263c0 1.597-1.78 2.55-3.11 1.664l-2.945-1.963l.888-1.332l-.888-1.333l2.946-1.963Zm-1.337 3.296L18 12.383V9.12l-2.447 1.632Zm1.361 10.406a1 1 0 0 1-1.32.507L12.5 20.29v.962a1 1 0 1 1-2 0v-.962l-3.094 1.375a1 1 0 1 1-.812-1.827l3.906-1.736v-3.85a1 1 0 0 1 2 0v3.85l3.906 1.736a1 1 0 0 1 .508 1.32Z"/><path d="M7.5 8.752a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4Zm6 0h-4v4h4v-4Z"/><path d="M13.793 8.959a1 1 0 0 1 0-1.415l3-3a1 1 0 1 1 1.414 1.415l-3 3a1 1 0 0 1-1.414 0Zm4.414 8a1 1 0 0 1-1.414 0l-3-3a1 1 0 0 1 1.414-1.415l3 3a1 1 0 0 1 0 1.415Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}