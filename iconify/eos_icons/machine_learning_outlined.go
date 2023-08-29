package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MachineLearningOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10a16.84 16.84 0 0 0-9 3.244A16.84 16.84 0 0 0 3 10v2.96a2.004 2.004 0 0 0-2 2.007v1.004c0 1.109 2 2.208 2 2.208v2.007a14.868 14.868 0 0 1 7.417 2.55A15.09 15.09 0 0 1 12 24a15.09 15.09 0 0 1 1.583-1.264A14.868 14.868 0 0 1 21 20.186v-2.208a2.004 2.004 0 0 0 2-2.007v-1.004a2.004 2.004 0 0 0-2-2.007Zm-9 11.422a16.841 16.841 0 0 0-7-2.996v-6.15a14.8 14.8 0 0 1 5.417 2.282A15.09 15.09 0 0 1 12 15.822a15.09 15.09 0 0 1 1.583-1.264A14.8 14.8 0 0 1 19 12.275v6.151a16.841 16.841 0 0 0-7 2.996ZM11 8h2v1h-2zm0-4h2v1h-2z"/><path fill="currentColor" d="M11 10h2v1h-2zM9 5a1 1 0 0 0 1-1a.983.983 0 0 0-.99-.99A.995.995 0 1 0 9 5Z"/><circle cx="15" cy="4" r="1" fill="currentColor"/><path fill="currentColor" d="M16 8H8a3.003 3.003 0 0 1-3-3V3a3.003 3.003 0 0 1 3-3h8a3.003 3.003 0 0 1 3 3v2a3.003 3.003 0 0 1-3 3ZM8 2a1.001 1.001 0 0 0-1 1v2a1.001 1.001 0 0 0 1 1h8a1.001 1.001 0 0 0 1-1V3a1.001 1.001 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}