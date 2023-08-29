package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.1 26.634l4.596 4.606m-6.429-14.036H22.1c2.598 0 4.705 2.11 4.705 4.715s-2.107 4.715-4.705 4.715m-6.809-2.41H9.5m5.791 0a3.51 3.51 0 0 1 0 7.02H9.5v-14.04h5.791a3.51 3.51 0 0 1 0 7.02Zm22.342-.572v-6.895l-3.701 4.631H38.5m-9.908-2.348c0-1.41 1.277-2.52 2.737-2.24c.957.184 1.722 1.012 1.817 1.983c.071.722-.157 1.434-.656 1.872c-.924.81-3.898 2.996-3.898 2.996h4.567"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}