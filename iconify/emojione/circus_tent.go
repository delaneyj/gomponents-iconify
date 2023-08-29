package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircusTent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M56.2 38H7.8L2 64h60z"/><path fill="#f2f2f2" d="M10.2 64h13.4l1.7-26H13.7z"/><path fill="#c94747" d="M2 38c0 2.3 2.6 4.2 5.8 4.2s5.8-1.9 5.8-4.2H2"/><path fill="#d0d0d0" d="M13.7 38c0 2.3 2.6 4.2 5.8 4.2c3.2 0 5.8-1.9 5.8-4.2H13.7"/><path fill="#f2f2f2" d="M53.8 64H40.4l-1.7-26h11.6z"/><path fill="#c94747" d="M62 38c0 2.3-2.6 4.2-5.8 4.2s-5.8-1.9-5.8-4.2H62"/><path fill="#d0d0d0" d="M50.3 38c0 2.3-2.6 4.2-5.8 4.2s-5.8-1.9-5.8-4.2h11.6"/><path fill="#ed4c5c" d="M32 12C30.8 22.4 17.8 38 2 38h60c-15.8 0-28.8-15.6-30-26"/><path fill="#fff" d="M13.7 38h11.6S32 30.1 32 12c0 0-6.7 26-18.3 26M32 12c0 18.1 6.7 26 6.7 26h11.6C38.7 38 32 12 32 12z"/><path fill="#c94747" d="M32 38h-6.7c0 2.3 3.4 4.2 6.7 4.2c3.2 0 6.7-1.9 6.7-4.2H32"/><path fill="#3e4347" d="M32 42.1L21.4 64h21.3z"/><path fill="#c94747" d="M21.4 64s-2.4 0-4.1-4.9c0 0 4.1-.5 14.7-16.9c0-.1-5.3 21.8-10.6 21.8m21.2 0s2.4 0 4.1-4.9c0 0-4.1-.5-14.7-16.9c0-.1 5.3 21.8 10.6 21.8"/><path fill="#94989b" d="M31.2 4.9h1.6v9.7h-1.6z"/><path fill="#d0d0d0" d="M31.6 4.9h.8v9.7h-.8z"/><ellipse cx="32" cy="5.4" fill="#ed4c5c" rx="1.5" ry="1.4"/><path fill="#c94747" d="M32 5.9c-.7 0-1.2-.4-1.5-.9c-.1.1-.1.3-.1.4c0 .8.7 1.4 1.5 1.4s1.5-.6 1.5-1.4c0-.1 0-.3-.1-.4c0 .5-.6.9-1.3.9"/><circle cx="32.4" cy="4.9" r=".4" fill="#ffc7ce"/><path fill="#ed4c5c" d="M32.8 7.1v4.1c2.9-3.1 5.8 1.1 8.6-2c-2.8 1.7-5.7-3.9-8.6-2.1"/>`),
		g.Group(children),
	)
}