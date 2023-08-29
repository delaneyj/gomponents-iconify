package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OschadTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="6.4" d="M5.5 38h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.436 19.65a1.943 1.943 0 0 1-1.936-1.938v-1.26c0-1.067.871-1.94 1.936-1.94s1.937.873 1.937 1.94v1.26c0 1.163-.872 1.938-1.937 1.938Zm7.57-5.137v5.148h-3.204v-5.148m6.408 0v5.148h-3.203m3.203-.001h.68v1.166m6.286-3.231a1.99 1.99 0 0 1-1.985 1.978h-.69c-.692 0-1.284-.593-1.284-1.285s.592-1.286 1.283-1.286v.005h2.676"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.691 14.995c.484-.484.775-.484 1.743-.484c1.064 0 1.742.485 1.742 1.648v3.489m7.324.974v-.97h-5.045v.97m4.365-.97V14.51h-3.007c0 2.134-.194 3.492-.873 5.14m-2.313 6.726l-1.581 6.481m-11.56-4.149c0-1.055.945-1.912 2.016-1.714c.693.131 1.26.79 1.323 1.516c.062.527-.126 1.12-.504 1.45c-.693.594-2.835 2.308-2.835 2.308h3.339m4.702 0v-5.293l-2.678 3.529h3.3m7.999 1.764l2.672-5.282h-3.293"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}