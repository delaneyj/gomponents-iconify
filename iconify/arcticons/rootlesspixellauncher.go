package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rootlesspixellauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.666 21.848L26.021 7.625a3.003 3.003 0 0 0-3.904 0L5.372 21.848a2.432 2.432 0 0 0-.305 3.426l.014.017a2.392 2.392 0 0 0 1.922.871h2.092v13.303a1.631 1.631 0 0 0 1.631 1.631h7.357a1.631 1.631 0 0 0 1.632-1.631h0v-8.188a1 1 0 0 1 1-1.001h6.647a1.001 1.001 0 0 1 1.001 1v8.189a1.631 1.631 0 0 0 1.632 1.631h7.357a1.631 1.631 0 0 0 1.631-1.631h0V26.162h2.052a2.442 2.442 0 0 0 1.591-4.314Z"/>`),
		g.Group(children),
	)
}