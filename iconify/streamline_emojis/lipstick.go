package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lipstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M11.44 44.36a12.56 1.64 0 1 0 25.12 0a12.56 1.64 0 1 0-25.12 0Z" opacity=".15"/><path fill="#87898c" d="m15.56 24.457l13.149-3.883a.73.73 0 0 1 .906.494l5.064 17.148a1.38 1.38 0 0 1-.933 1.714l-11.902 3.514a1.38 1.38 0 0 1-1.714-.932l-5.067-17.158a.73.73 0 0 1 .494-.907Z"/><path fill="#bdbec0" d="m34.67 38.22l-5.06-17.15a.73.73 0 0 0-.9-.5l-2.32.69l5.66 19.17l1.69-.5a1.37 1.37 0 0 0 .93-1.71Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m15.56 24.457l13.149-3.883a.73.73 0 0 1 .906.494l5.064 17.148a1.38 1.38 0 0 1-.933 1.714l-11.902 3.514a1.38 1.38 0 0 1-1.714-.932l-5.067-17.158a.73.73 0 0 1 .494-.907Z"/><path fill="#daedf7" d="m28.74 20.56l-13.22 3.91l-.26-2.94a.68.68 0 0 1 .49-.72l10.82-3.2a.69.69 0 0 1 .8.34Z"/><path fill="#e8f4fa" d="M23.62 18.49L25 21.68l3.76-1.12L27.37 18a.69.69 0 0 0-.8-.34l-2.95.87Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m28.74 20.56l-13.22 3.91l-.26-2.94a.68.68 0 0 1 .49-.72l10.82-3.2a.69.69 0 0 1 .8.34Z"/><path fill="#ff6242" d="m25.79 17.84l-9.26 2.74l-4.86-16.46a1.39 1.39 0 0 1 1.84-1.67l6.9 2.75a4.14 4.14 0 0 1 2.43 2.67Z"/><path fill="#ff866e" d="m12.23 6l6 2.39a4.14 4.14 0 0 1 2.43 2.67l2.25 7.61l2.87-.85l-2.95-10a4.14 4.14 0 0 0-2.42-2.62l-6.9-2.75a1.39 1.39 0 0 0-1.84 1.67Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m25.79 17.84l-9.26 2.74l-4.86-16.46a1.39 1.39 0 0 1 1.84-1.67l6.9 2.75a4.14 4.14 0 0 1 2.43 2.67Z"/><path fill="#ff6242" d="m17.987 35.236l14.54-4.293l.977 3.309l-14.54 4.293z"/><path fill="#ff866e" d="m33.5 34.26l-3.01.88l-.97-3.3l3.01-.89l.97 3.31z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m17.987 35.236l14.54-4.293l.977 3.309l-14.54 4.293z"/>`),
		g.Group(children),
	)
}