package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegularShapePt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M47.268 41.5v7.768H39.5v5.464h7.768V62.5h5.464v-7.768H60.5v-5.464h-7.768V41.5Z" color="currentColor"/><path fill="currentColor" d="M50.1 11.502a3.5 3.5 0 0 0-2.157.666L11.141 38.906a3.5 3.5 0 0 0-1.27 3.912l14.057 43.264a3.5 3.5 0 0 0 3.328 2.418h45.488a3.5 3.5 0 0 0 3.328-2.418l10.475-32.24c.3.022.598.045.902.045c6.89 0 12.551-5.661 12.551-12.551c0-6.89-5.661-12.549-12.55-12.549c-3.141 0-6.018 1.185-8.227 3.117L52.057 12.168a3.5 3.5 0 0 0-1.957-.666zm-.1 7.824l25.422 18.47a12.394 12.394 0 0 0-.522 3.54c0 4.1 2.012 7.753 5.086 10.049L70.201 81.5H29.8L17.313 43.074L50 19.326zm37.45 16.461c3.106 0 5.55 2.442 5.55 5.549s-2.444 5.55-5.55 5.55c-3.107 0-5.55-2.443-5.55-5.55c0-3.107 2.443-5.549 5.55-5.549z" color="currentColor"/>`),
		g.Group(children),
	)
}