package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiltersOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.75a5.25 5.25 0 1 0 0 10.5a5.25 5.25 0 0 0 0-10.5ZM5.25 8a6.75 6.75 0 1 1 13.274 1.738A6.752 6.752 0 1 1 12 21.438a6.75 6.75 0 1 1-6.524-11.7A6.76 6.76 0 0 1 5.25 8Zm.77 3.136A5.252 5.252 0 0 0 8 21.25a5.25 5.25 0 0 0 5.079-6.586a6.75 6.75 0 0 1-7.058-3.529Zm8.504 3.126c.148.555.226 1.138.226 1.738a6.723 6.723 0 0 1-1.625 4.393a5.25 5.25 0 1 0 4.854-9.258a6.775 6.775 0 0 1-3.455 3.127Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}