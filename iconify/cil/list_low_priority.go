package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListLowPriority(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 56h240v32H256zm0 106.667h240v32H256zm0 106.666h240v32H256zM328 376h168v32H328z"/><path fill="currentColor" d="M161.231 408h5.969v68.783L302 393l-134.8-86.228V376h-5.965C98.8 376 48 311.4 48 232S98.8 88 161.231 88H216V56h-54.769C121.783 56 84.91 74.755 57.4 108.81C30.7 141.866 16 185.616 16 232s14.7 90.134 41.4 123.19C84.91 389.245 121.783 408 161.231 408Zm37.969-42.772l42.8 27.381l-42.8 26.608Z"/>`),
		g.Group(children),
	)
}