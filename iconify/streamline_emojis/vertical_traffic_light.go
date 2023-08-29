package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalTrafficLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M13 46.5a11 1.5 0 1 0 22 0a11 1.5 0 1 0-22 0Z" opacity=".15"/><path fill="#daedf7" d="M24 45.5A10.47 10.47 0 0 1 13.53 35V11A10.47 10.47 0 0 1 24 .5A10.47 10.47 0 0 1 34.47 11v24A10.47 10.47 0 0 1 24 45.5Z"/><path fill="#fff" d="M24 .5A10.47 10.47 0 0 0 13.53 11v2.61a10.47 10.47 0 0 1 20.94 0V11A10.47 10.47 0 0 0 24 .5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 45.5h0A10.47 10.47 0 0 1 13.53 35V11A10.47 10.47 0 0 1 24 .5h0A10.47 10.47 0 0 1 34.47 11v24A10.47 10.47 0 0 1 24 45.5Z"/><path fill="#87898c" d="M24 42.88A7.85 7.85 0 0 1 16.15 35V11A7.85 7.85 0 0 1 24 3.12A7.85 7.85 0 0 1 31.85 11v24A7.85 7.85 0 0 1 24 42.88Z"/><path fill="#656769" d="M24 3.12A7.85 7.85 0 0 0 16.15 11v2.75a7.85 7.85 0 0 1 15.7 0V11A7.85 7.85 0 0 0 24 3.12Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 42.88h0A7.85 7.85 0 0 1 16.15 35V11A7.85 7.85 0 0 1 24 3.12h0A7.85 7.85 0 0 1 31.85 11v24A7.85 7.85 0 0 1 24 42.88Z"/><path fill="#ffe500" d="M19.5 23.5a4.5 4.5 0 1 0 9 0a4.5 4.5 0 1 0-9 0Z"/><path fill="#fff48c" d="M24 19a4.49 4.49 0 0 0-1.92 8.56l5-7.31A4.49 4.49 0 0 0 24 19Z"/><path fill="#6dd627" d="M19.5 12a4.5 4.5 0 1 0 9 0a4.5 4.5 0 1 0-9 0Z"/><path fill="#9ceb60" d="M24 7.5a4.49 4.49 0 0 0-1.92 8.56l5-7.31A4.49 4.49 0 0 0 24 7.5Z"/><path fill="#ff6242" d="M19.5 35a4.5 4.5 0 1 0 9 0a4.5 4.5 0 1 0-9 0Z"/><path fill="#ff866e" d="M24 30.5a4.49 4.49 0 0 0-1.92 8.56l5-7.31A4.49 4.49 0 0 0 24 30.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.5 23.5a4.5 4.5 0 1 0 9 0a4.5 4.5 0 1 0-9 0Zm0-11.5a4.5 4.5 0 1 0 9 0a4.5 4.5 0 1 0-9 0Zm0 23a4.5 4.5 0 1 0 9 0a4.5 4.5 0 1 0-9 0Z"/><path fill="#bdbec0" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M29 34.7c0 2.76-.81-2.5-5-2.5s-5 5.26-5 2.5a5 5 0 0 1 10 0Zm0-11.5c0 2.76-.81-2.5-5-2.5s-5 5.3-5 2.5a5 5 0 0 1 10 0Zm0-11.5c0 2.76-.81-2.5-5-2.5s-5 5.26-5 2.5a5 5 0 0 1 10 0Z"/>`),
		g.Group(children),
	)
}