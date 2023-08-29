package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackplayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.46 22.39v1.36m-13.64-5.97v3.7a.6.6 0 0 0 .34.54l3 1.44a.6.6 0 0 1 0 1.08l-3 1.44a.6.6 0 0 0-.34.54v3.7a1.2 1.2 0 0 0 1.72 1.08l14.16-6.76a.6.6 0 0 0 0-1.08L19.54 16.7a1.2 1.2 0 0 0-1.72 1.08Zm8.11 1.97v4m2.76 0v-2.68"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.86 2.52a21.48 21.48 0 0 0-22 17.79a1 1 0 0 0 .56 1.06l4.37 2.09a.6.6 0 0 1 0 1.08L3.4 26.63a1 1 0 0 0-.56 1.06a21.49 21.49 0 1 0 22-25.17Z"/>`),
		g.Group(children),
	)
}