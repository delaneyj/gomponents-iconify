package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoMobilePhones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 16c0 7.732-6.268 14-14 14S2 23.732 2 16S8.268 2 16 2s14 6.268 14 14Zm-5.233 6.645A10.952 10.952 0 0 0 27 16c0-6.075-4.925-11-11-11c-2.497 0-4.8.832-6.645 2.233l1.652 1.652A1 1 0 0 1 12 8h8a1 1 0 0 1 1 1v9.879l3.767 3.766Zm-2.122 2.122l-1.651-1.652A1 1 0 0 1 20 24h-8a1 1 0 0 1-1-1v-9.879L7.233 9.355A10.952 10.952 0 0 0 5 16c0 6.075 4.925 11 11 11c2.497 0 4.8-.832 6.645-2.233ZM12 14.12v7.629c0 .138.112.25.25.25h7.5a.248.248 0 0 0 .105-.023L12 14.12Zm8 3.758V11.25a.25.25 0 0 0-.25-.25h-6.629L20 17.879ZM13 9.5a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 0-1h-5a.5.5 0 0 0-.5.5Z"/>`),
		g.Group(children),
	)
}