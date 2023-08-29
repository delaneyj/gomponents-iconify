package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Metro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="256" cy="256" r="256" fill="#2B3B47"/><path fill="#FF473E" d="M96 64h320v368H96z"/><path fill="#F9F9F7" d="M104 72h304v224H104z"/><path fill="#FFD469" d="M207 340c0 11.046-8.954 20-20 20s-20-8.954-20-20s8.954-20 20-20s20 8.954 20 20zm116-20c-11.046 0-20 8.954-20 20s8.954 20 20 20s20-8.954 20-20s-8.954-20-20-20z"/><path fill="#E5E4DF" d="M192.12 432h32l-14.854 59.414a238.256 238.256 0 0 1-57.386-19.175L192.12 432zM320 432h-32l14.854 59.414a238.256 238.256 0 0 0 57.386-19.175L320 432z"/><path fill="#132028" d="M396 280H116V136h280v144zM353 88H161v32h192V88zm-33 312H192v32h128v-32z"/>`),
		g.Group(children),
	)
}