package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugDroidBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m197.35 43.62l11.14-11.13a12 12 0 0 0-17-17l-12.15 12.2a91.84 91.84 0 0 0-102.68 0L64.49 15.51a12 12 0 0 0-17 17l11.16 11.11A91.58 91.58 0 0 0 36 104v40a92 92 0 0 0 184 0v-40a91.58 91.58 0 0 0-22.65-60.38ZM196 104v4H60v-4a68 68 0 0 1 136 0Zm-68 108a68.07 68.07 0 0 1-68-68v-12h136v12a68.07 68.07 0 0 1-68 68Zm12-132a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm-56 0a16 16 0 1 1 16 16a16 16 0 0 1-16-16Z"/>`),
		g.Group(children),
	)
}