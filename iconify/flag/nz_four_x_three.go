package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NzFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><g id="flagNz4x30"><g id="flagNz4x31"><path d="M0-.3v.5l1-.5z"/><path d="M.2.3L0-.1l1-.2z"/></g><use href="#flagNz4x31" transform="scale(-1 1)"/><use href="#flagNz4x31" transform="rotate(72 0 0)"/><use href="#flagNz4x31" transform="rotate(-72 0 0)"/><use href="#flagNz4x31" transform="scale(-1 1) rotate(72)"/></g></defs><path fill="#00247d" fill-rule="evenodd" d="M0 0h640v480H0z"/><g transform="translate(-111 36.1) scale(.66825)"><use width="100%" height="100%" fill="#fff" href="#flagNz4x30" transform="matrix(45.4 0 0 45.4 900 120)"/><use width="100%" height="100%" fill="#cc142b" href="#flagNz4x30" transform="matrix(30 0 0 30 900 120)"/></g><g transform="rotate(82 525.2 114.6) scale(.66825)"><use width="100%" height="100%" fill="#fff" href="#flagNz4x30" transform="rotate(-82 519 -457.7) scale(40.4)"/><use width="100%" height="100%" fill="#cc142b" href="#flagNz4x30" transform="rotate(-82 519 -457.7) scale(25)"/></g><g transform="rotate(82 525.2 114.6) scale(.66825)"><use width="100%" height="100%" fill="#fff" href="#flagNz4x30" transform="rotate(-82 668.6 -327.7) scale(45.4)"/><use width="100%" height="100%" fill="#cc142b" href="#flagNz4x30" transform="rotate(-82 668.6 -327.7) scale(30)"/></g><g transform="translate(-111 36.1) scale(.66825)"><use width="100%" height="100%" fill="#fff" href="#flagNz4x30" transform="matrix(50.4 0 0 50.4 900 480)"/><use width="100%" height="100%" fill="#cc142b" href="#flagNz4x30" transform="matrix(35 0 0 35 900 480)"/></g><path fill="#012169" d="M0 0h320v240H0z"/><path fill="#fff" d="m37.5 0l122 90.5L281 0h39v31l-120 89.5l120 89V240h-40l-120-89.5L40.5 240H0v-30l119.5-89L0 32V0z"/><path fill="#c8102e" d="M212 140.5L320 220v20l-135.5-99.5zm-92 10l3 17.5l-96 72H0zM320 0v1.5l-124.5 94l1-22L295 0zM0 0l119.5 88h-30L0 21z"/><path fill="#fff" d="M120.5 0v240h80V0zM0 80v80h320V80z"/><path fill="#c8102e" d="M0 96.5v48h320v-48zM136.5 0v240h48V0z"/>`),
		g.Group(children),
	)
}