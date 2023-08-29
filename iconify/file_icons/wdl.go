package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wdl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M83.772 192.028L.838 143.648V47.636L84.634 0l82.934 48.379v68.338l66.21 41.31l-35.18 22.616l-45.15-28.428l-69.676 39.813zM0 367.608v96.014L82.935 512l83.796-47.635v-96.013l-82.935-48.38L0 367.609zm272.261-205.109l-83.685 47.635v71.464L127.4 318.013l34.343 23.453l41.841-26.5l67.818 39.56l83.795-47.635v-96.013L272.261 162.5z"/>`),
		g.Group(children),
	)
}