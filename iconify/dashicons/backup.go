package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.65 2.88c3.93 2.01 5.48 6.84 3.47 10.77s-6.83 5.48-10.77 3.47a7.942 7.942 0 0 1-3.86-4.4l1.64-1.03a6.125 6.125 0 0 0 3.08 3.76c3.01 1.54 6.69.35 8.23-2.66A6.114 6.114 0 1 0 4.56 7.21l1.88.97l-4.95 3.08l-.39-5.82l1.78.91C4.9 2.4 9.75.89 13.65 2.88zm-4.36 7.83A.997.997 0 0 1 9 10c0-.07.03-.12.04-.19h-.01L10 5l.97 4.81L14 13l-4.5-2.12l.02-.02c-.08-.04-.16-.09-.23-.15z"/>`),
		g.Group(children),
	)
}