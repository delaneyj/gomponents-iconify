package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M241.72 113a11.88 11.88 0 0 0-9.73-5H212V88a12 12 0 0 0-12-12h-69.34a4 4 0 0 1-2.4-.8l-27.73-20.8a12.05 12.05 0 0 0-7.2-2.4H40a12 12 0 0 0-12 12v144a4 4 0 0 0 4 4h179.09a4 4 0 0 0 3.79-2.74l28.49-85.47a11.86 11.86 0 0 0-1.65-10.79ZM40 60h53.33a4 4 0 0 1 2.4.8l27.73 20.8a12.07 12.07 0 0 0 7.2 2.4H200a4 4 0 0 1 4 4v20H69.76a12 12 0 0 0-11.38 8.21L36 183.35V64a4 4 0 0 1 4-4Zm195.78 61.26L208.2 204H37.55L66 118.74a4 4 0 0 1 3.76-2.74H232a4 4 0 0 1 3.79 5.26Z"/>`),
		g.Group(children),
	)
}