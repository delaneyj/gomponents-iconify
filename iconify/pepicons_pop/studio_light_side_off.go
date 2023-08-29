package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioLightSideOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.89 4.456C15.22 3.57 17 4.523 17 6.12v3.263c0 1.597-1.78 2.55-3.11 1.664l-2.945-1.963l.888-1.332l-.888-1.333l2.946-1.963Zm-1.337 3.296L15 9.383V6.12l-2.447 1.632Zm1.361 10.406a1 1 0 0 1-1.32.507L9.5 17.29v.962a1 1 0 1 1-2 0v-.962l-3.094 1.375a1 1 0 1 1-.812-1.827L7.5 15.102v-3.85a1 1 0 0 1 2 0v3.85l3.906 1.736a1 1 0 0 1 .508 1.32Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4.5 5.752a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4Zm6 0h-4v4h4v-4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.793 5.959a1 1 0 0 1 0-1.415l3-3a1 1 0 1 1 1.414 1.415l-3 3a1 1 0 0 1-1.414 0Zm4.414 8a1 1 0 0 1-1.414 0l-3-3a1 1 0 0 1 1.414-1.415l3 3a1 1 0 0 1 0 1.415Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}