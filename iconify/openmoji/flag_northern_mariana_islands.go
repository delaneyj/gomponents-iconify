package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagNorthernMarianaIslands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61b2e4" d="M5 17h62v38H5z"/><ellipse cx="36" cy="36" fill="none" stroke="#fff" stroke-miterlimit="10" stroke-width="3" rx="10.466" ry="10.5"/><path fill="none" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="M46.466 35.931a10.466 10.466 0 1 1-20.932.138"/><path fill="none" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="M25.534 36.069a10.466 10.466 0 1 1 20.932-.138"/><path fill="#9b9b9a" stroke="#9b9b9a" stroke-linecap="round" stroke-linejoin="round" d="M40.5 50h-9l3-18h3l3 18z"/><path fill="#9b9b9a" stroke="#9b9b9a" stroke-linecap="round" stroke-linejoin="round" d="M38.5 29h-5l1 4h3l1-4z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m32.122 42l3.983-12l3.434 11.816L30 34.696l12-.296l-9.878 7.6z"/><path fill="none" stroke="#5c9e31" stroke-linecap="round" stroke-linejoin="round" d="M25.747 38.116A10.642 10.642 0 0 1 25.724 34m20.53 0a10.639 10.639 0 0 1 .023 4.116"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}