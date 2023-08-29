package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WinkingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M.54 20.38a18.88 18.88 0 1 0 37.76 0a18.88 18.88 0 1 0-37.76 0Z"/><path fill="#ebcb00" d="M19.42 1.5A18.88 18.88 0 1 0 38.3 20.38A18.88 18.88 0 0 0 19.42 1.5Zm0 34.92A17.23 17.23 0 1 1 36.64 19.2a17.23 17.23 0 0 1-17.22 17.22Z"/><path fill="#fff48c" d="M13.76 5.28a5.66 1.42 0 1 0 11.32 0a5.66 1.42 0 1 0-11.32 0Z"/><path fill="#45413c" d="M4.32 44.58a15.1 1.42 0 1 0 30.2 0a15.1 1.42 0 1 0-30.2 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M.54 20.38a18.88 18.88 0 1 0 37.76 0a18.88 18.88 0 1 0-37.76 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" d="M24.14 27.91a8.5 8.5 0 0 1-9.44 0"/><path fill="#ffaa54" d="M33.11 25.1c0 .78-1.06 1.41-2.36 1.41s-2.36-.63-2.36-1.41s1.05-1.42 2.36-1.42s2.36.63 2.36 1.42Zm-27.38 0c0 .78 1.06 1.41 2.36 1.41s2.36-.63 2.36-1.41s-1-1.42-2.36-1.42s-2.36.63-2.36 1.42Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m29.8 18.96l-1.89 2.36l4.25-.94"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.51 20.38a.95.95 0 0 1-1.89 0a.95.95 0 1 1 1.89 0Z"/><path fill="#fff" stroke="#45413c" stroke-linejoin="round" d="M47.24 20.81a.59.59 0 0 1-.27.43a.55.55 0 0 1-.5 0l-4.89-1.91a.56.56 0 0 1 .14-1.09l5.14-.61a.57.57 0 0 1 .47.16a.6.6 0 0 1 .17.46Zm-.65-10.26a.55.55 0 0 0-.35-.36a.58.58 0 0 0-.51.06l-4.36 2.91a.57.57 0 0 0 .37 1l5.15-.51a.57.57 0 0 0 .43-.25a.61.61 0 0 0 .06-.49ZM40.78 4a.56.56 0 0 0-.5-.07a.57.57 0 0 0-.36.35l-1.72 5a.57.57 0 0 0 .23.66a.56.56 0 0 0 .7-.06l3.79-3.52a.53.53 0 0 0 .18-.46a.54.54 0 0 0-.24-.42Z"/>`),
		g.Group(children),
	)
}