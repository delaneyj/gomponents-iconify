package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotFaceOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M4.5 45.5a19.5 1.5 0 1 0 39 0a19.5 1.5 0 1 0-39 0Z" opacity=".15"/><path fill="#daedf7" d="M42 40.5a4 4 0 0 1-4 4H10a4 4 0 0 1-4-4V27a18 18 0 0 1 36 0Z"/><path fill="#fff" d="M24 9A18 18 0 0 0 6 27v5a18 18 0 0 1 36 0v-5A18 18 0 0 0 24 9Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M42 40.5a4 4 0 0 1-4 4H10a4 4 0 0 1-4-4V27a18 18 0 0 1 36 0Z"/><path fill="#c0dceb" d="M17.5 34h13v12h-13Z"/><path fill="#adc4d9" d="M17.5 38.5h13v2h-13z"/><path fill="#adc4d9" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45.5 35.92a1 1 0 0 1-.5.87l-3 1.71V26l3 1.71a1 1 0 0 1 .5.87Zm-43 0a1 1 0 0 0 .5.87l3 1.71V26l-3 1.71a1 1 0 0 0-.5.87Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 9a9 9 0 0 1 10.5-6.73"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.5 3.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.5 34h13v12h-13Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.5 34h9a2 2 0 0 1 2 2v2.5h0h-13h0V36a2 2 0 0 1 2-2Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.5 11.71a18 18 0 0 0-19 0V24h19Z"/><path fill="#00dfeb" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M27.5 25a5 5 0 1 0 10 0a5 5 0 1 0-10 0Z"/><path fill="#627b8c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30.5 25a2 2 0 1 0 4 0a2 2 0 1 0-4 0Z"/><path fill="#00dfeb" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.5 25a5 5 0 1 0 10 0a5 5 0 1 0-10 0Z"/><path fill="#627b8c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13.5 25a2 2 0 1 0 4 0a2 2 0 1 0-4 0Z"/>`),
		g.Group(children),
	)
}