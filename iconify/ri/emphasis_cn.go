package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmphasisCn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm-5.5 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm11 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3ZM13 2v2h6v2h-1.968a18.222 18.222 0 0 1-3.621 6.302a14.683 14.683 0 0 0 5.327 3.042l-.536 1.93A16.686 16.686 0 0 1 12 13.726a16.697 16.697 0 0 1-6.2 3.547l-.536-1.929a14.7 14.7 0 0 0 5.327-3.042a18.077 18.077 0 0 1-2.822-4.3h2.24A16.03 16.03 0 0 0 12 10.876A16.17 16.17 0 0 0 14.91 6H5V4h6V2h2Z"/>`),
		g.Group(children),
	)
}