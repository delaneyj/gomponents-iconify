package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KrFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagKr4x30"><path fill-opacity=".7" d="M-95.8-.4h682.7v512H-95.8z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagKr4x30)" transform="translate(89.8 .4) scale(.9375)"><path fill="#fff" d="M-95.8-.4H587v512H-95.8Z"/><g transform="rotate(-56.3 361.6 -101.3) scale(10.66667)"><g id="flagKr4x31"><path id="flagKr4x32" fill="#000" d="M-6-26H6v2H-6Zm0 3H6v2H-6Zm0 3H6v2H-6Z"/><use width="100%" height="100%" y="44" href="#flagKr4x32"/></g><path stroke="#fff" d="M0 17v10"/><path fill="#cd2e3a" d="M0-12a12 12 0 0 1 0 24Z"/><path fill="#0047a0" d="M0-12a12 12 0 0 0 0 24A6 6 0 0 0 0 0Z"/><circle cy="-6" r="6" fill="#cd2e3a"/></g><g transform="rotate(-123.7 191.2 62.2) scale(10.66667)"><use width="100%" height="100%" href="#flagKr4x31"/><path stroke="#fff" d="M0-23.5v3M0 17v3.5m0 3v3"/></g></g>`),
		g.Group(children),
	)
}