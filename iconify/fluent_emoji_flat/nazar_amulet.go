package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NazarAmulet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs><path id="fluentEmojiFlatNazarAmulet0" fill="#fff" d="M16.2 23.8a7.68 7.68 0 1 0 0-15.36a7.68 7.68 0 0 0 0 15.36Z"/></defs><g fill="none"><path fill="#1345B7" d="M16 30c7.732 0 14-6.268 14-14S23.732 2 16 2S2 8.268 2 16s6.268 14 14 14Z"/><use href="#fluentEmojiFlatNazarAmulet0"/><use href="#fluentEmojiFlatNazarAmulet0"/><path fill="#26C9FC" d="M16.2 21.24a5.12 5.12 0 1 0 0-10.24a5.12 5.12 0 0 0 0 10.24Z"/><path fill="#212121" d="M16.2 18.68a2.56 2.56 0 1 0 0-5.12a2.56 2.56 0 0 0 0 5.12Z"/></g>`),
		g.Group(children),
	)
}