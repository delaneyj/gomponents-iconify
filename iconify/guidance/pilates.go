package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pilates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 20.5h24m-10-11H3.5a2 2 0 0 0-1.831 1.195c-.214.485.042 1.016.417 1.39l4.53 4.531a3.018 3.018 0 0 0 4.268 0L15 12.5h2s3.036.174 6.5-1.826M10 14L5.5 9.5m-2.805-2S.885 6.943.56 5.724a1.768 1.768 0 0 1 1.242-2.163a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61l-.291.079Z"/>`),
		g.Group(children),
	)
}