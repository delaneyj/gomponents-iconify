package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Turtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.39 4.5v9.28L16 17.51l-8-4.65Zm3.22 0l14.47 8.35L32 17.5l-6.43-3.73ZM6.31 15.64l8 4.64v7.43l-8 4.65Zm35.38 0v16.71l-8-4.64v-7.43ZM24 16.57l6.43 3.71v7.43L24 31.43l-6.43-3.72v-7.43Zm-8 13.92l6.43 3.73v9.28L7.92 35.14Zm16 0l8 4.65l-14.39 8.36v-9.28Z"/>`),
		g.Group(children),
	)
}