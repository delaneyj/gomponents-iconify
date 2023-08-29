package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M576 768V384H320v256q0 53 37.5 90.5T448 768h128zm1024 448v192H448v-192l128-192H448q-159 0-271.5-112.5T64 640V320L0 256l32-128h480L544 0h960l32 192l-64 32v800z"/>`),
		g.Group(children),
	)
}