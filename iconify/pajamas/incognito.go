package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Incognito(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.248 2.042a.75.75 0 0 0-.933.403L2.513 6.5H.75a.75.75 0 1 0 0 1.5h2.233a.76.76 0 0 0 .033 0H15.25a.75.75 0 0 0 0-1.5h-1.763l-1.802-4.055a.75.75 0 0 0-.933-.403L8 3.005l-2.752-.963ZM11.846 6.5l-1.25-2.814l-2.348.822a.75.75 0 0 1-.496 0l-2.347-.822L4.155 6.5h7.69Zm-6.346 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 1.5A2.5 2.5 0 0 0 8 11.5a2.5 2.5 0 0 0 5 0h.25a.75.75 0 0 0 0-1.5h-.75a2.496 2.496 0 0 0-2-1c-.818 0-1.544.393-2 1h-1a2.496 2.496 0 0 0-2-1c-.818 0-1.544.393-2 1h-.75a.75.75 0 0 0 0 1.5H3A2.5 2.5 0 0 0 5.5 14Zm5-1.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}