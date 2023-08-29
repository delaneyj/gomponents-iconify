package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontDownloadOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.775 22.575l-.6-.575H4q-.825 0-1.413-.588T2 20V4.825L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3ZM4 20h13.175L4 6.825V20Zm18-.875l-2-2V4H6.875l-2-2H20q.825 0 1.413.588T22 4v15.125Zm-6.4-6.4l-3.05-3.05l-.5-1.425h-.1l-.225.6l-1.35-1.35l.325-.875q.125-.275.375-.45t.55-.175h.75q.3 0 .55.175t.375.45l2.3 6.1Zm-5 .675Zm2.85-2.85Zm.625 3.5l3.075 3.075q-.05.35-.312.613t-.663.262q-.3 0-.55-.175t-.35-.45l-1.2-3.325Zm-4.475.9l-.875 2.425q-.1.275-.35.45t-.55.175q-.5 0-.812-.412t-.113-.913l2.7-7.1l1.4 1.4l-.8 2.225h3.025l1.75 1.75H9.6Z"/>`),
		g.Group(children),
	)
}