package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 301V88q0-27 12.5-44.5t38-26t53-11.5t67-3t67 3t53 11.5t38 26T341 88v213q0 28-21 48v38q0 8-6.5 14.5T299 408h-22q-8 0-14.5-6.5T256 387v-22H85v22q0 8-6 14.5T64 408H43q-9 0-15.5-6.5T21 387v-38Q0 329 0 301zm74.5 22q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5zm192 0q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5zM299 195V88H43v107h256z"/>`),
		g.Group(children),
	)
}