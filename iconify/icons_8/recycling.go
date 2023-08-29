package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Recycling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3.97a3.264 3.264 0 0 0-2.75 1.5l-2.813 4.624l1.72 1.03l2.812-4.593c.54-.805 1.563-.823 2.092 0l3.125 5l-1.5.876L23.313 15V9.687l-1.437.844l-3.125-5.06a3.264 3.264 0 0 0-2.75-1.5zm-5.594 8.124l-4.5 2.594l1.25.75l-2.562 4.218l-.032-.03c-.012.017.012.044 0 .06C3.092 21.916 4.788 25 7.5 25H13v-2H7.5c-1.267 0-1.946-1.25-1.28-2.22l.03-.03v-.03l2.625-4.25l1.53.936v-5.312zm14.656 3.562l-1.718 1.063l2.5 4c.64 1.087-.087 2.28-1.25 2.28H19v-2l-4.906 3L19 27v-2h5.594c2.637 0 4.328-3 2.968-5.313v-.03l-2.5-4z"/>`),
		g.Group(children),
	)
}