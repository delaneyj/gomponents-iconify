package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ansible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m14.156 15.297l6.25 4.927l-4.141-10.214zM16 0C7.161 0 0 7.161 0 16s7.161 16 16 16s16-7.161 16-16S24.839 0 16 0zm7.729 23.073a1.136 1.136 0 0 1-1.167 1.109c-.313 0-.552-.12-.885-.391l-8.255-6.667l-2.771 6.938H8.255L15.25 7.255a1.054 1.054 0 0 1 1.021-.677c.432-.016.839.25.99.677l6.365 15.323c.057.151.104.313.104.464v.031z"/>`),
		g.Group(children),
	)
}