package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmFrames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#BBDDF5" d="M29 11h6v14h-6zM1 11h6v14H1zm8 0h18v14H9z"/><path fill="#231F20" d="M35 11V6H1v5h6v14H1v5h34v-5h-6V11h6zm-8 14H9V11h18v14z"/><path fill="#F5F8FA" d="M31 7h2v3h-2zm-4 0h2v3h-2zm-4 0h2v3h-2zm-4 0h2v3h-2zm-4 0h2v3h-2zm-4 0h2v3h-2zM7 7h2v3H7zM3 7h2v3H3zm28 19h2v3h-2zm-4 0h2v3h-2zm-4 0h2v3h-2zm-4 0h2v3h-2zm-4 0h2v3h-2zm-4 0h2v3h-2zm-4 0h2v3H7zm-4 0h2v3H3z"/><path fill="#88C9F9" d="M29 25V11h1v14zM9 25V11h1v14z"/>`),
		g.Group(children),
	)
}