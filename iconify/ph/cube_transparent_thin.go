package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeTransparentThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M219.92 95.22a1.8 1.8 0 0 0-.1-.33a3.38 3.38 0 0 0-.13-.42a2.58 2.58 0 0 0-.19-.37c-.06-.11-.11-.22-.17-.32a3.62 3.62 0 0 0-.5-.6l-56-56a2.32 2.32 0 0 0-.33-.27A4 4 0 0 0 160 36H40a4 4 0 0 0-4 4v120a4 4 0 0 0 .9 2.5a2.32 2.32 0 0 0 .27.33l56 56A4 4 0 0 0 96 220h120a4 4 0 0 0 4-4V96a4.13 4.13 0 0 0-.08-.78ZM164 49.66L206.34 92H164ZM92 206.34L49.66 164H92ZM92 156H44V49.66l48 48ZM49.66 44H156v48H97.66ZM156 100v56h-56v-56Zm-56 112v-48h58.34l48 48Zm112-5.66l-48-48V100h48Z"/>`),
		g.Group(children),
	)
}