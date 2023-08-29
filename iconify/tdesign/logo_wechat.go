package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoWechat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.796 17.027H8.75c-1.153 0-2.254-.188-3.262-.53L2.65 17.92l.352-2.712C1.162 13.855 0 11.861 0 9.64c0-4.083 3.918-7.39 8.75-7.39c4.174 0 7.665 2.468 8.54 5.77a9.394 9.394 0 0 0-.6-.02c-4.364 0-8.19 3.037-8.19 7.11c0 .67.104 1.312.296 1.917ZM6 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm5.5.007a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path fill="currentColor" d="M21.874 19.52C23.187 18.405 24 16.863 24 15.16C24 11.758 20.754 9 16.75 9S9.5 11.758 9.5 15.161c0 3.403 3.246 6.161 7.25 6.161c.95 0 1.856-.155 2.686-.437l2.438 1.407V19.52Zm-7.564-5.362a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm4.88 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}