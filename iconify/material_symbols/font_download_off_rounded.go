package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontDownloadOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.775 22.575l-7.65-7.625l1.95-.9l6.875 6.875q-.25.5-.725.788t-1.05.287H4q-.825 0-1.413-.588T2 20V4.825L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3ZM22 19.125l-6.4-6.4l-2.3-6.1q-.125-.275-.375-.45T12.375 6h-.75q-.3 0-.55.175t-.375.45l-.325.875l-5.5-5.5H20q.825 0 1.413.588T22 4v15.125ZM9.6 14.95h2.525l-1.85-1.85L8.8 11.625l-1.9 5.05q-.2.475.1.9t.8.425q.325 0 .563-.175t.362-.475l.875-2.4Zm2.35-6.7h.1l.5 1.425l-.825-.825l.225-.6Zm2.125 5.8l1.2 3.325q.1.275.35.45t.55.175q.4 0 .663-.263t.312-.612l-3.075-3.075Z"/>`),
		g.Group(children),
	)
}