package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForLesotho(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9f9f9" d="M2 32c0 5.9 1.7 11.4 4.6 16h50.7c2.9-4.6 4.6-10.1 4.6-16s-1.7-11.4-4.6-16H6.6C3.7 20.6 2 26.1 2 32z"/><path fill="#428bc1" d="M57.4 16C52.1 7.6 42.7 2 32 2S11.9 7.6 6.6 16h50.8z"/><path fill="#83bf4f" d="M6.6 48c5.3 8.4 14.7 14 25.4 14s20.1-5.6 25.4-14H6.6z"/><path fill="#3e4347" d="M35.1 30.8c.9-1.1 1.4-2.7 1.4-4.5c0-3.5-2-6.2-4.5-6.2s-4.5 2.7-4.5 6.2c0 1.8.5 3.4 1.4 4.5l-6.9 8s0 5.3 10 5.3s10-5.3 10-5.3l-6.9-8m-.4-4.6c0 1.2-.3 2.3-.8 3.1l-1.5-1.7V26l1.4-.7l-1.4-.7v-2.8c1.3.4 2.3 2.3 2.3 4.4m-5.4 0c0-2.2 1-4 2.3-4.4v2.8l-1.4.7l1.4.7v1.6l-1.5 1.7c-.5-.8-.8-1.9-.8-3.1"/>`),
		g.Group(children),
	)
}