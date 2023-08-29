package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatsCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M229.66 187.13a78 78 0 0 0-61.5-112.71a78 78 0 1 0-141.82 64.71l-7.87 27.57a12 12 0 0 0 14.83 14.83l27.57-7.87a78.25 78.25 0 0 0 26.94 7.9a78.05 78.05 0 0 0 107.32 40.1l27.57 7.87a12 12 0 0 0 14.83-14.83Zm-168.13-25.9a5.82 5.82 0 0 0-1.65.23L30 170l8.53-29.87a6 6 0 0 0-.5-4.53A66 66 0 1 1 64.41 162a6.1 6.1 0 0 0-2.88-.77Zm155.93 26.89L226 218l-29.87-8.53a6 6 0 0 0-4.53.5a66 66 0 0 1-90.48-28.15a77.92 77.92 0 0 0 71-94.68A66 66 0 0 1 218 183.59a6 6 0 0 0-.54 4.53Z"/>`),
		g.Group(children),
	)
}