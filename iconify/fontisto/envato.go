package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.781.274c-.88-.471-2.922-.327-5.324.383c-1.825 1.472-7.76 6.766-7.835 13.348a.49.49 0 0 1-.94.195l-.001-.003c-.566-1.316-1.026-3.512-.405-6.888a.492.492 0 0 0-.863-.401l-.001.001q-.324.4-.614.813c-3.787 5.469-1.079 12.48 3.378 14.914s11.081 1.894 14.272-3.943s.555-17.226-1.667-18.419z"/>`),
		g.Group(children),
	)
}