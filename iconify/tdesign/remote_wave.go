package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.004 3a9.991 9.991 0 0 0-8.426 4.612l-.54.841l-1.684-1.079l.54-.842A11.991 11.991 0 0 1 12.004 1c4.247 0 7.978 2.207 10.11 5.532l.539.842l-1.684 1.08l-.54-.842A9.992 9.992 0 0 0 12.004 3Zm-.012 3.988A4.995 4.995 0 0 0 7.78 9.294l-.54.842l-1.683-1.079l.54-.842a6.995 6.995 0 0 1 5.896-3.227a6.995 6.995 0 0 1 5.897 3.227l.54.842l-1.684 1.08l-.54-.843a4.995 4.995 0 0 0-4.213-2.306ZM5 11h14v12h-2V13H7v10H5V11Zm6 4h2.004v2.004H11V15Z"/>`),
		g.Group(children),
	)
}