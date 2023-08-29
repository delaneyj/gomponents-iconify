package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShootingStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f5a300" d="m7 31l14.1 10.3L14.2 62l17.9-12.8l18 12.8L57 2z"/><path id="emojioneShootingStar0" fill="#fff" d="m38.9 39.9l10.9-7.8H36.3l-4.2-12.6L28 32.1H14.5l10.9 7.8l-4.1 12.6l10.8-7.8L43 52.5z"/><use href="#emojioneShootingStar0"/><path fill="#ffce31" d="M39 28.4h3.5L57 2L37.8 24.7zM57 2L34.6 15l1.8 5.6zm-3.1 26.4L57 2l-9.8 26.4z"/>`),
		g.Group(children),
	)
}