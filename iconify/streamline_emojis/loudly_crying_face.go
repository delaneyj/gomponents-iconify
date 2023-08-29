package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoudlyCryingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17 21.5H8.5m22.5 0h8.5"/><path fill="#ffb0ca" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.94 29.5a.94.94 0 0 0-.71.33a1 1 0 0 0-.22.76a7.09 7.09 0 0 0 14 0a1 1 0 0 0-.22-.76a.94.94 0 0 0-.71-.33Z"/><path fill="#ff87af" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M29.57 33.84A11.37 11.37 0 0 0 24 32.53a11.37 11.37 0 0 0-5.57 1.31a7.16 7.16 0 0 0 11.14 0Z"/><path fill="#4aeff7" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.5 36.26a20 20 0 0 0 5 3.35V21.5h-5Zm27 0a20 20 0 0 1-5 3.35V21.5h5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/>`),
		g.Group(children),
	)
}