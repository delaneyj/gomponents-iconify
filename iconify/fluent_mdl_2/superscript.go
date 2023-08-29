package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1701 486q-6 6-10 12t-10 14h239v128h-384v-64q0-52 19-94t47-77t62-64t61-54t48-48t19-47q0-26-19-45t-45-19q-26 0-45 19t-19 45h-128q0-40 15-75t41-61t61-41t75-15q40 0 75 15t61 41t41 61t15 75q0 38-15 73t-41 63h-1l-162 158zm-261 26l-384 512l384 512h-160l-304-405l-304 405H512l384-512l-384-512h160l304 405l304-405h160z"/>`),
		g.Group(children),
	)
}