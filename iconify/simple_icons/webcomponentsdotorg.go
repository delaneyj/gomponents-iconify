package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webcomponentsdotorg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.731 2.225l-.01.016H5.618L0 11.979l5.618 9.736h12.8l.04.06l2.134-3.735l.518-.893h-.008l.008-.014l-.626-.75h.895l.006-.01l.008.01L24 11.994l-2.607-4.39l-.003.005l-.011-.02h-.945l.63-.763l-2.606-4.57l-.006.01l-.024-.04H11.73zM9.107 6.824h6.19l-.53.764h-.023l2.398 4.015h.875l-.277.328l.357.435h-.956l-2.398 4.015h.027l.523.764H9.074l-2.99-5.168l3.022-5.155z"/>`),
		g.Group(children),
	)
}