package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaLink2Fill0"><g id="evaLink2Fill1"><g id="evaLink2Fill2" fill="currentColor"><path d="m13.29 9.29l-4 4a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0-1.42-1.42Z"/><path d="M12.28 17.4L11 18.67a4.2 4.2 0 0 1-5.58.4a4 4 0 0 1-.27-5.93l1.42-1.43a1 1 0 0 0 0-1.42a1 1 0 0 0-1.42 0l-1.27 1.28a6.15 6.15 0 0 0-.67 8.07a6.06 6.06 0 0 0 9.07.6l1.42-1.42a1 1 0 0 0-1.42-1.42Zm7.38-14.18a6.18 6.18 0 0 0-8.13.68L10.45 5a1.09 1.09 0 0 0-.17 1.61a1 1 0 0 0 1.42 0L13 5.3a4.17 4.17 0 0 1 5.57-.4a4 4 0 0 1 .27 5.95l-1.42 1.43a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.42-1.42a6.06 6.06 0 0 0-.6-9.06Z"/></g></g></g>`),
		g.Group(children),
	)
}