package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassDone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M8.91 45.19a15.09 1.81 0 1 0 30.18 0a15.09 1.81 0 1 0-30.18 0Z" opacity=".15"/><path fill="#80ddff" d="M34.86 8.28H13.14l.35 5.72a7.26 7.26 0 0 0 3.68 5.86l1.84 1a2.83 2.83 0 0 1 0 4.93l-1.84 1a7.26 7.26 0 0 0-3.68 5.86l-.35 5.72h21.72l-.35-5.72a7.26 7.26 0 0 0-3.68-5.86l-1.84-1a2.83 2.83 0 0 1 0-4.93l1.84-1A7.26 7.26 0 0 0 34.51 14Z"/><path fill="#b8ecff" d="M30.83 26.84L29 25.8a2.83 2.83 0 0 1 0-4.93l1.84-1A7.26 7.26 0 0 0 34.51 14l.35-5.73h-7.54l-9.15 30.19h9.34L31 26.94Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.86 8.28H13.14l.35 5.72a7.26 7.26 0 0 0 3.68 5.86l1.84 1a2.83 2.83 0 0 1 0 4.93l-1.84 1a7.26 7.26 0 0 0-3.68 5.86l-.35 5.72h21.72l-.35-5.72a7.26 7.26 0 0 0-3.68-5.86l-1.84-1a2.83 2.83 0 0 1 0-4.93l1.84-1A7.26 7.26 0 0 0 34.51 14Z"/><path fill="#debb7e" d="M11.32 2.25h25.35v6.04H11.32Z"/><path fill="#f0d5a8" d="M35.47 2.25H12.53a1.21 1.21 0 0 0-1.21 1.2V6a1.21 1.21 0 0 1 1.21-1.21h22.94A1.21 1.21 0 0 1 36.68 6V3.45a1.21 1.21 0 0 0-1.21-1.2Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M11.32 2.25h25.35v6.04H11.32Z"/><path fill="#debb7e" d="M11.32 38.46h25.35v6.04H11.32Z"/><path fill="#f0d5a8" d="M35.47 38.46H12.53a1.21 1.21 0 0 0-1.21 1.21v2.57A1.21 1.21 0 0 1 12.53 41h22.94a1.21 1.21 0 0 1 1.21 1.21v-2.54a1.21 1.21 0 0 0-1.21-1.21Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M11.32 38.46h25.35v6.04H11.32Z"/><path fill="#f7e5c6" d="M34.44 32.13c-1-1-4-3.06-10.44-3.06s-9.52 2.38-10.49 3.41a2.28 2.28 0 0 1 0 .26l-.35 5.72h21.7l-.35-5.72c-.02-.21-.04-.41-.07-.61Z"/><path fill="#fff5e3" d="m21 29.26l-2.78 9.2h9.34l2.61-8.61a21.92 21.92 0 0 0-6.17-.78a23.4 23.4 0 0 0-3 .19Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.44 32.13c-1-1-4-3.06-10.44-3.06s-9.52 2.38-10.49 3.41a2.28 2.28 0 0 1 0 .26l-.35 5.72h21.7l-.35-5.72c-.02-.21-.04-.41-.07-.61Z"/>`),
		g.Group(children),
	)
}